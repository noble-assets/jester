package keys

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/cosmos/go-bip39"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"jester.noble.xyz/v2/appstate"
)

// setupTestKeyring creates a temporary keyring for testing
func setupTestKeyring(t *testing.T) (*appstate.AppState, func()) {
	tempDir, err := os.MkdirTemp("", "jester-test-*")
	require.NoError(t, err)

	appState := &appstate.AppState{
		Home: tempDir,
	}

	appState.NewCodec()

	err = appState.NewKeyring()
	require.NoError(t, err)

	cleanup := func() {
		os.RemoveAll(tempDir)
	}

	return appState, cleanup
}

func TestCreateMnemonic(t *testing.T) {
	mnemonic, err := createMnemonic()
	require.NoError(t, err)

	words := strings.Split(mnemonic, " ")
	assert.Equal(t, 24, len(words), "mnemonic should be 24 words")

	assert.True(t, bip39.IsMnemonicValid(mnemonic), "mnemonic should be valid BIP39")
}

func TestKeyCreation(t *testing.T) {
	appState, cleanup := setupTestKeyring(t)
	defer cleanup()

	cmd := createCmd(appState)
	err := cmd.RunE(cmd, []string{})
	require.NoError(t, err)

	assert.True(t, appState.KeyExists())
}
func TestKeyDeletion(t *testing.T) {
	appState, cleanup := setupTestKeyring(t)
	defer cleanup()

	createCmd := createCmd(appState)
	err := createCmd.RunE(createCmd, []string{})
	require.NoError(t, err)

	// Setup delete command with mock yes input
	deleteCmd := deleteCmd(appState)
	deleteCmd.SetIn(bytes.NewBufferString("y\n"))

	err = deleteCmd.RunE(deleteCmd, []string{})
	require.NoError(t, err)

	assert.False(t, appState.KeyExists())
}

func TestKeyRecovery(t *testing.T) {
	appState, cleanup := setupTestKeyring(t)
	defer cleanup()

	mnemonic := "bind guilt bid science stage slight loan soft strong asthma vacant train prize vast envelope rocket race praise flavor blush text elder nothing public"
	expectedAddress := "jester1gr2ru8wpgr4ltjnqxfrld64kqc0ljqhnwsc6qa"

	recoverCmd := recoverCmd(appState)
	recoverCmd.SetIn(bytes.NewBufferString(mnemonic + "\n"))

	var output bytes.Buffer
	recoverCmd.SetOut(&output)

	err := recoverCmd.RunE(recoverCmd, []string{})
	require.NoError(t, err)

	assert.True(t, appState.KeyExists())

	outputStr := output.String()
	assert.Contains(t, outputStr, expectedAddress, "recover command output should contain the expected address")
}

func TestShowKey(t *testing.T) {
	appState, cleanup := setupTestKeyring(t)
	defer cleanup()

	cmd := createCmd(appState)
	err := cmd.RunE(cmd, []string{})
	require.NoError(t, err)

	var output bytes.Buffer
	showCmd := showCmd(appState)
	showCmd.SetOut(&output)

	err = showCmd.RunE(showCmd, []string{})
	require.NoError(t, err)

	outputStr := output.String()
	assert.Contains(t, outputStr, "name:")
	assert.Contains(t, outputStr, "type:")
	assert.Contains(t, outputStr, "address:")
	assert.Contains(t, outputStr, "pubkey:")
}
