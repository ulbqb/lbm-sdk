package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTokenType(t *testing.T) {
	tokenType := NewBaseTokenType(defaultContractID, defaultTokenType, defaultName)

	require.Equal(t, `{"contract_id":"abcdef01","token_type":"10000001","name":"name"}`, tokenType.String())

	var tokenType2 TokenType
	bz, err := ModuleCdc.MarshalJSON(tokenType)
	require.NoError(t, err)
	err = ModuleCdc.UnmarshalJSON(bz, &tokenType2)
	require.NoError(t, err)

	require.Equal(t, defaultName, tokenType2.GetName())
	require.Equal(t, defaultContractID, tokenType2.GetContractID())
	require.Equal(t, defaultTokenType, tokenType2.GetTokenType())

	require.Equal(t, tokenType.GetName(), tokenType2.GetName())
	require.Equal(t, tokenType.GetContractID(), tokenType2.GetContractID())
	require.Equal(t, tokenType.GetTokenType(), tokenType2.GetTokenType())

	require.Equal(t, `{"contract_id":"abcdef01","token_type":"10000001","name":"name"}`, tokenType.String())

	tokenType3 := tokenType.SetName("testname")
	require.Equal(t, defaultName, tokenType.GetName())
	require.Equal(t, "testname", tokenType3.GetName())
}
