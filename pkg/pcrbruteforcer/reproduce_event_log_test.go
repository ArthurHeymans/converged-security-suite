package pcrbruteforcer

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/9elements/converged-security-suite/v2/pkg/pcr"
	"github.com/9elements/converged-security-suite/v2/pkg/registers"
	"github.com/9elements/converged-security-suite/v2/pkg/tpmeventlog"
	"github.com/9elements/converged-security-suite/v2/pkg/uefi"
	"github.com/9elements/converged-security-suite/v2/testdata/firmware"
	"github.com/google/go-tpm/tpm2"
	"github.com/stretchr/testify/require"
)

func unhex(fataler fataler, h string) []byte {
	b, err := hex.DecodeString(h)
	if err != nil {
		fataler.Fatal(err)
	}
	return b
}

func TestReproduceEventLog(t *testing.T) {
	firmwareImage := firmware.FakeIntelFirmware

	firmware, err := uefi.ParseUEFIFirmwareBytes(firmwareImage)
	require.NoError(t, err)

	measureOptions := []pcr.MeasureOption{
		pcr.SetFlow(pcr.FlowIntelCBnT0T),
		pcr.SetIBBHashDigest(tpm2.AlgSHA1),
		pcr.SetRegisters(registers.Registers{
			registers.ParseACMPolicyStatusRegister(0x0000000200108682),
		}),
	}
	measurements, _, debugInfo, err := pcr.GetMeasurements(firmware, 0, measureOptions...)
	require.NoError(t, err, fmt.Sprintf("debugInfo: '%v'", debugInfo))

	eventLog := &tpmeventlog.TPMEventLog{
		Events: []*tpmeventlog.Event{
			{
				PCRIndex: 0,
				Type:     tpmeventlog.EV_NO_ACTION,
				Data:     []byte("StartupLocality\000\003"),
				Digest: &tpmeventlog.Digest{
					HashAlgo: 4,
					Digest:   unhex(t, "0000000000000000000000000000000000000000"),
				},
			},
			{
				PCRIndex: 0,
				Type:     tpmeventlog.EV_S_CRTM_CONTENTS,
				Digest: &tpmeventlog.Digest{
					HashAlgo: 4,
					Digest:   unhex(t, "527C9A38B2F45FBF89C382547E0A0812722A47D3"),
				},
			},
			{
				PCRIndex: 0,
				Type:     tpmeventlog.EV_S_CRTM_VERSION,
				Digest: &tpmeventlog.Digest{
					HashAlgo: 4,
					Digest:   unhex(t, "C14F556E35C9BB45F189B03F383A6A3E31256681"),
				},
			},
			{
				PCRIndex: 0,
				Type:     tpmeventlog.EV_POST_CODE,
				Digest: &tpmeventlog.Digest{
					HashAlgo: 4,
					Digest:   unhex(t, "4C9836F73CC42ADBECE7D565B783E618B4A75C22"),
				},
			},
			{
				PCRIndex: 0,
				Type:     tpmeventlog.EV_SEPARATOR,
				Digest: &tpmeventlog.Digest{
					HashAlgo: 4,
					Digest:   unhex(t, "9069CA78E7450A285173431B3E52C5C25299E473"),
				},
			},
		},
	}

	succeeded, acmPolicyStatus, _ := ReproduceEventLog(eventLog, measurements, firmwareImage)
	require.True(t, succeeded)
	require.NotNil(t, acmPolicyStatus)
	require.Equal(t, uint64(0x0000000200108681), acmPolicyStatus.Raw())
}
