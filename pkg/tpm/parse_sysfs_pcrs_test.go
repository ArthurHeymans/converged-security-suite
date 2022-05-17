package tpm

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseSysfsPCRS(t *testing.T) {
	sample := `PCR-00: 09 7E 00 E3 B8 D9 A8 EA 07 CE B5 30 92 06 0A AC 2D 16 60 F5
PCR-01: 27 8B 8D EE 2F EA AA 43 B2 3C F4 88 93 36 AD F5 04 EE 93 D8
PCR-02: E6 E1 13 9A 4C 17 72 0F 13 BE DD D1 1D 29 5B 8B FF D5 D7 E8
PCR-03: B2 A8 3B 0E BF 2F 83 74 29 9A 5B 2B DF C3 1E A9 55 AD 72 36
PCR-04: 0D 43 13 B3 B4 7A 97 AB 21 EC 45 CE 93 2A 4B 86 54 D6 61 19
PCR-05: B2 A8 3B 0E BF 2F 83 74 29 9A 5B 2B DF C3 1E A9 55 AD 72 36
PCR-06: B2 A8 3B 0E BF 2F 83 74 29 9A 5B 2B DF C3 1E A9 55 AD 72 36
PCR-07: 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
PCR-08: 89 2B 6C B7 10 80 D8 0E AF D2 31 48 BF 91 D1 3B 4C BF 0D 4E
PCR-09: 3E 4F 1B AB 5E 7B 17 1A 49 BA E3 D1 49 43 9E 89 CC B2 EE 7D
PCR-10: 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
PCR-11: 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
PCR-12: 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
PCR-13: 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
PCR-14: 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
PCR-15: 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
PCR-16: 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
PCR-17: FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF
PCR-18: FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF
PCR-19: FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF
PCR-20: FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF
PCR-21: FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF
PCR-22: FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF
PCR-23: 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
`
	pcrs, err := parseSysfsPCRs([]byte(sample))
	require.NoError(t, err)

	hexDecode := func(in string) []byte {
		result, err := hex.DecodeString(strings.ReplaceAll(in, " ", ""))
		require.NoError(t, err)
		return result
	}
	require.Equal(t, hexDecode(`09 7E 00 E3 B8 D9 A8 EA 07 CE B5 30 92 06 0A AC 2D 16 60 F5`), pcrs[0])
	require.Equal(t, hexDecode(`FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF FF`), pcrs[22])
	require.Equal(t, hexDecode(`00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00`), pcrs[23])
}
