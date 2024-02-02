package blockinfo

import (
	"math/big"
	"testing"

	"github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/core/types"
	ethTypes "github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/smt/pkg/smt"
)

/*
Contrived tests of the SMT inserts (compared with test cases from the JS implementation)
*/
func TestBlockInfoHeader(t *testing.T) {
	tests := []struct {
		BlockHash          string
		CoinbaseAddress    string
		NewBlockNumber     uint64
		BlockGasLimit      uint64
		FinalTimestamp     uint64
		FinalGER           string
		L1BlochHash        string
		FinalBlockInfoRoot string
	}{
		{
			BlockHash:          "0x1fe466d9df83e1d2a4c32e21c6078b8f5f590e7db30b006965faa2f27a9b4fea",
			CoinbaseAddress:    "0x617b3a3528F9cDd6630fd3301B9c8911F7Bf063D",
			NewBlockNumber:     1,
			BlockGasLimit:      4294967295,
			FinalTimestamp:     1944498031,
			FinalGER:           "0x0000000000000000000000000000000000000000000000000000000000000000",
			L1BlochHash:        "0x0000000000000000000000000000000000000000000000000000000000000000",
			FinalBlockInfoRoot: "0x64f37448decfd2837a23f825060a705b1135247a08f88a047ccde3aa58efb52b",
		}, {
			BlockHash:          "0x4a9bfcb163ec91c5beb22e6aca41592433092c8c7821b01d37fd0de483f9265d",
			CoinbaseAddress:    "0x617b3a3528F9cDd6630fd3301B9c8911F7Bf063D",
			NewBlockNumber:     1,
			BlockGasLimit:      4294967295,
			FinalTimestamp:     1944498031,
			FinalGER:           "0x0000000000000000000000000000000000000000000000000000000000000000",
			L1BlochHash:        "0x0000000000000000000000000000000000000000000000000000000000000000",
			FinalBlockInfoRoot: "0x445c76b4a370754cd2fbb52da85e9c5265e9a10244ebf751f0de27a252efe4a0",
		}, {
			BlockHash:          "0x4a9bfcb163ec91c5beb22e6aca41592433092c8c7821b01d37fd0de483f9265d",
			CoinbaseAddress:    "0x617b3a3528F9cDd6630fd3301B9c8911F7Bf063D",
			NewBlockNumber:     5,
			BlockGasLimit:      4294967295,
			FinalTimestamp:     1944498031,
			FinalGER:           "0x0000000000000000000000000000000000000000000000000000000000000000",
			L1BlochHash:        "0x0000000000000000000000000000000000000000000000000000000000000000",
			FinalBlockInfoRoot: "0xf8c8d52e97e83cbe07ad1883f6510ec2aafcde26e5d291290ecd240e76241bce",
		},
	}

	for _, test := range tests {
		infoTree := NewBlockInfoTree()
		blockHash := common.HexToHash(test.BlockHash)
		coinbaseAddress := common.HexToAddress(test.CoinbaseAddress)
		ger := common.HexToHash(test.FinalGER)
		l1BlochHash := common.HexToHash(test.L1BlochHash)

		err := infoTree.InitBlockHeader(
			&blockHash,
			&coinbaseAddress,
			test.NewBlockNumber,
			test.BlockGasLimit,
			test.FinalTimestamp,
			&ger,
			&l1BlochHash,
		)
		if err != nil {
			t.Fatal(err)
		}

		root := common.BigToHash(infoTree.GetRoot()).Hex()

		if root != test.FinalBlockInfoRoot {
			t.Fatalf("expected root %s, got %s", test.FinalBlockInfoRoot, root)
		}
	}
}

func TestSetBlockTx(t *testing.T) {
	tests := []struct {
		txIndex             int
		receipt             ethTypes.Receipt
		logIndex            int64
		cumulativeGasUsed   uint64
		effectivePercentage uint8
		finalBlockInfoRoot  string
	}{
		{
			txIndex: 0,
			receipt: ethTypes.Receipt{
				Status: 1,
				TxHash: common.HexToHash("0xd2a69c7d3c99953bae9d273f2375b2653bd4ec47a05eefc5fc7c041d07fba8a0"),
				Logs: []*types.Log{
					{
						Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Topics: []common.Hash{
							common.HexToHash("0000000000000000000000000000000000000000000000000000000000000001"),
						},
						Data: common.HexToHash("0000000000000000000000000000000000000000000000000000000000000000").Bytes(),
					},
					{
						Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Topics: []common.Hash{
							common.HexToHash("0000000000000000000000000000000000000000000000000000000000000004"),
							common.HexToHash("0000000000000000000000000000000000000000000000000000000000000005"),
							common.HexToHash("0000000000000000000000000000000000000000000000000000000000000006"),
						},
						Data: common.HexToHash("0000000000000000000000000000000000000000000000000000000000000000").Bytes(),
					},
					{
						Address: common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Topics: []common.Hash{
							common.HexToHash("0000000000000000000000000000000000000000000000000000000000000001"),
							common.HexToHash("0000000000000000000000000000000000000000000000000000000000000002"),
						},
						Data: common.HexToHash("0000000000000000000000000000000000000000000000000000000000000000").Bytes(),
					},
				},
			},
			logIndex:            0,
			cumulativeGasUsed:   26336,
			effectivePercentage: 255,
			finalBlockInfoRoot:  "0x763711586b99a8c51ddfb765ff17ad8beb9312b246126d84ff02cea3cfc39828",
		}, {
			txIndex: 0,
			receipt: ethTypes.Receipt{
				Status: 0,
				TxHash: common.HexToHash("0xac65e2fd657a4ee6318cc66cf98b05ae74ce3f0f3982370af951176e7b599c2c"),
				Logs:   []*types.Log{},
			},
			logIndex:            0,
			cumulativeGasUsed:   21000,
			effectivePercentage: 0,
			finalBlockInfoRoot:  "0xeb85acdbf2dd2d0c9b2637124520b09816065596a2e6f8d8869ffd22850371e4",
		}, {
			txIndex: 0,
			receipt: ethTypes.Receipt{
				Status: 1,
				TxHash: common.HexToHash("0xac65e2fd657a4ee6318cc66cf98b05ae74ce3f0f3982370af951176e7b599c2c"),
				Logs:   []*types.Log{},
			},
			logIndex:            0,
			cumulativeGasUsed:   21000,
			effectivePercentage: 255,
			finalBlockInfoRoot:  "0x4d94273d5028e71194fc29e4e73e1d94a392e65e5cc64ca86a80b3811fd18f51",
		}, {
			txIndex: 0,
			receipt: ethTypes.Receipt{
				Status: 1,
				TxHash: common.HexToHash("0x8f9b0375a6b0f1bd9d54ff499921766828ae8e5314fc44a494736b5c4cc3bb56"),
				Logs:   []*types.Log{},
			},
			logIndex:            0,
			cumulativeGasUsed:   10000,
			effectivePercentage: 255,
			finalBlockInfoRoot:  "0x42e4e630df29444b15bf553e03354cb4bac013220ca0cea2ea88e8f5efc26131",
		},
	}

	for _, test := range tests {
		infoTree := NewBlockInfoTree()

		root, err := infoTree.SetBlockTx(
			test.txIndex,
			&test.receipt,
			test.logIndex,
			test.cumulativeGasUsed,
			test.effectivePercentage,
		)
		if err != nil {
			t.Fatal(err)
		}
		rootHex := common.BigToHash(root).Hex()

		if rootHex != test.finalBlockInfoRoot {
			t.Fatalf("expected root %s, got %s", test.finalBlockInfoRoot, rootHex)
		}
	}
}

func TestBlockComulativeGasUsed(t *testing.T) {
	tests := []struct {
		gasUsed      uint64
		expectedRoot string
	}{
		{
			gasUsed:      26336,
			expectedRoot: "0x5cd280355924dcf29ac41ccae98d678091d182af191443f3c92562e1c1c64254",
		}, {
			gasUsed:      21000,
			expectedRoot: "0x9cfdda40abe9331804fe6b55be89421bd74ca56e9da719e39bbf5518e08155e1",
		}, {
			gasUsed:      10000,
			expectedRoot: "0x32cc19445bc8843c9f432cad24c3c6ea198734547d996bb977a2011c04d917f8",
		},
	}

	for i, test := range tests {
		infoTree := NewBlockInfoTree()

		root, err := infoTree.SetBlockGasUsed(test.gasUsed)
		if err != nil {
			t.Fatal(err)
		}
		actualRoot := common.BigToHash(root).Hex()
		// root taken from JS implementation
		if actualRoot != test.expectedRoot {
			t.Fatalf("Test %d expected root %s, got %s", i+1, test.expectedRoot, actualRoot)
		}
	}
}

func TestSetL2BlockHash(t *testing.T) {
	tests := []struct {
		blockHash    string
		expectedRoot string
	}{
		{
			blockHash:    "0x1fe466d9df83e1d2a4c32e21c6078b8f5f590e7db30b006965faa2f27a9b4fea",
			expectedRoot: "0x1db6a2e2ce5016d114c38a4530c66adfb1b24bf66714d20eb983ed4910ed6600",
		},
		{
			blockHash:    "0x4a9bfcb163ec91c5beb22e6aca41592433092c8c7821b01d37fd0de483f9265d",
			expectedRoot: "0xaa99d2be4188527344ef32d31024b127006e9fbbdb75862de564d448c47816be",
		},
	}

	for i, test := range tests {
		smt := smt.NewSMT(nil)
		blockHash := common.HexToHash(test.blockHash)

		root, err := setL2BlockHash(smt, &blockHash)
		if err != nil {
			t.Fatal(err)
		}
		actualRoot := common.BigToHash(root).Hex()
		// root taken from JS implementation
		if actualRoot != test.expectedRoot {
			t.Fatalf("Test %d expected root %s, got %s", i+1, test.expectedRoot, actualRoot)
		}
	}
}

func TestSetCoinbase(t *testing.T) {
	tests := []struct {
		coinbaseAddress string
		expectedRoot    string
	}{
		{
			coinbaseAddress: "0x617b3a3528F9cDd6630fd3301B9c8911F7Bf063D",
			expectedRoot:    "0x27fb3bd76956839741006a2dd73bfffadb9573c6cd8ce60b0566b7c81a55b7b4",
		},
	}

	for i, test := range tests {
		smt := smt.NewSMT(nil)
		coinbaseAddress := common.HexToAddress(test.coinbaseAddress)

		root, err := setCoinbase(smt, &coinbaseAddress)
		if err != nil {
			t.Fatal(err)
		}
		actualRoot := common.BigToHash(root).Hex()
		// root taken from JS implementation
		if actualRoot != test.expectedRoot {
			t.Fatalf("Test %d expected root %s, got %s", i+1, test.expectedRoot, actualRoot)
		}
	}
}

func TestSetBlockNumber(t *testing.T) {
	tests := []struct {
		blockNum     uint64
		expectedRoot string
	}{
		{
			blockNum:     1,
			expectedRoot: "0x45685d4b214d4eb330627ff12797a4063fefcc13579f5c1fe5f7131a397c26b4",
		}, {
			blockNum:     5,
			expectedRoot: "0xad832d8f6f2ca140d3aff0065d7fb920a643e3619ead5404832e54a511aeec6c",
		},
	}

	for i, test := range tests {
		smt := smt.NewSMT(nil)

		root, err := setBlockNumber(smt, test.blockNum)
		if err != nil {
			t.Fatal(err)
		}
		actualRoot := common.BigToHash(root).Hex()
		// root taken from JS implementation
		if actualRoot != test.expectedRoot {
			t.Fatalf("Test %d expected root %s, got %s", i+1, test.expectedRoot, actualRoot)
		}
	}
}

func TestSetGasLimit(t *testing.T) {
	tests := []struct {
		gasLimit     uint64
		expectedRoot string
	}{
		{
			gasLimit:     4294967295,
			expectedRoot: "0xdfb45af6d25ba1d98cf29e5272049fc5007d63fe4a0c0ca2322ef826debb2b6c",
		},
	}

	for i, test := range tests {
		smt := smt.NewSMT(nil)

		root, err := setGasLimit(smt, test.gasLimit)
		if err != nil {
			t.Fatal(err)
		}
		actualRoot := common.BigToHash(root).Hex()
		// root taken from JS implementation
		if actualRoot != test.expectedRoot {
			t.Fatalf("Test %d expected root %s, got %s", i+1, test.expectedRoot, actualRoot)
		}
	}
}

func TestSetTimestamp(t *testing.T) {
	tests := []struct {
		timestamp    uint64
		expectedRoot string
	}{
		{
			timestamp:    1944498031,
			expectedRoot: "0xe0ef08c2c9c75a9e7a9fceec0483414489be3b9d34312115a2eb9c30339a3922",
		},
	}

	for i, test := range tests {
		smt := smt.NewSMT(nil)

		root, err := setTimestamp(smt, test.timestamp)
		if err != nil {
			t.Fatal(err)
		}
		actualRoot := common.BigToHash(root).Hex()
		// root taken from JS implementation
		if actualRoot != test.expectedRoot {
			t.Fatalf("Test %d expected root %s, got %s", i+1, test.expectedRoot, actualRoot)
		}
	}
}

func TestSetGer(t *testing.T) {
	tests := []struct {
		ger          string
		expectedRoot string
	}{
		{
			ger:          "0x819feaf48e670e06a9faa2ecce4b795f214ed1f0258b22e49db7691da8206663",
			expectedRoot: "0x61f1fac06c5b64bf969df3e57cea7418fdab1c38e3ee5ac654b2c74e27316bd4",
		}, {
			ger:          "0xb15aa2b6ef32f2b517e19672e43186094f7e0d37a4b60b77644ee33b5feb3f7f",
			expectedRoot: "0xf598491f603545710aa7ec6ad8c9b2f554c0f02eb04092d992228e9dfcb682e0",
		}, {
			ger:          "0x5f4e0c5cbfc891af492d7335d988c2578204a75c997bfad0e7ca8fc2bd4389c9",
			expectedRoot: "0x7a0b0cc58dc3777704c34d965f6b5d86146280c82b288c23a32aee1989d1a504",
		},
	}

	for i, test := range tests {
		smt := smt.NewSMT(nil)
		ger := common.HexToHash(test.ger)

		root, err := setGer(smt, &ger)
		if err != nil {
			t.Fatal(err)
		}
		actualRoot := common.BigToHash(root).Hex()
		// root taken from JS implementation
		if actualRoot != test.expectedRoot {
			t.Fatalf("Test %d expected root %s, got %s", i+1, test.expectedRoot, actualRoot)
		}
	}
}

func TestSetL1BlockHash(t *testing.T) {
	tests := []struct {
		l1BlockHash  string
		expectedRoot string
	}{
		{
			l1BlockHash:  "0x819feaf48e670e06a9faa2ecce4b795f214ed1f0258b22e49db7691da8206663",
			expectedRoot: "0xc0cea75b3047bf5f28cf3affaeaf9842e68a5d29544a237e4e8bbea4b369d25f",
		}, {
			l1BlockHash:  "0xb15aa2b6ef32f2b517e19672e43186094f7e0d37a4b60b77644ee33b5feb3f7f",
			expectedRoot: "0x68909800f942475ab88aea079b7407131f7e1aad2de0a860803411f9560803a7",
		}, {
			l1BlockHash:  "0x5f4e0c5cbfc891af492d7335d988c2578204a75c997bfad0e7ca8fc2bd4389c9",
			expectedRoot: "0xcb2eb84e4e2070d4c7aa827ab796131339c20554a20592c0f80afa225a9e5901",
		},
	}

	for i, test := range tests {
		smt := smt.NewSMT(nil)
		l1BlockHash := common.HexToHash(test.l1BlockHash)

		root, err := setL1BlockHash(smt, &l1BlockHash)
		if err != nil {
			t.Fatal(err)
		}
		actualRoot := common.BigToHash(root).Hex()
		// root taken from JS implementation
		if actualRoot != test.expectedRoot {
			t.Fatalf("Test %d expected root %s, got %s", i+1, test.expectedRoot, actualRoot)
		}
	}
}

func TestSetL2TxHash(t *testing.T) {
	smt := smt.NewSMT(nil)
	txIndex := big.NewInt(1)
	l2TxHash := common.HexToHash("0x000000000000000000000000000000005Ca1aB1E").Big()

	root, err := setL2TxHash(smt, txIndex, l2TxHash)
	if err != nil {
		t.Fatal(err)
	}

	// root taken from JS implementation
	expectedRoot := "a9127a157cee3cd2452a194e4efc2f8a5612cfc36c66e768700727ede4d0e2e6"
	actualRoot := root.Text(16)

	if actualRoot != expectedRoot {
		t.Fatalf("expected root %s, got %s", expectedRoot, actualRoot)
	}
}

func TestSetTxStatus(t *testing.T) {
	smt := smt.NewSMT(nil)
	txIndex := big.NewInt(1)
	status := common.HexToHash("0x000000000000000000000000000000005Ca1aB1E").Big()

	root, err := setTxStatus(smt, txIndex, status)
	if err != nil {
		t.Fatal(err)
	}

	// root taken from JS implementation
	expectedRoot := "7cb6a0928f5165a422cfbe5f93d1cc9eda3f686715639823f6087818465fcbb8"
	actualRoot := root.Text(16)

	if actualRoot != expectedRoot {
		t.Fatalf("expected root %s, got %s", expectedRoot, actualRoot)
	}
}

func TestSetCumulativeGasUsed(t *testing.T) {
	smt := smt.NewSMT(nil)
	txIndex := big.NewInt(1)
	cgu := common.HexToHash("0x000000000000000000000000000000005Ca1aB1E").Big()

	root, err := setCumulativeGasUsed(smt, txIndex, cgu)
	if err != nil {
		t.Fatal(err)
	}

	// root taken from JS implementation
	expectedRoot := "c07ff46f07be5b81465c30848202acc4bf82805961d8a9f9ffe74e820e4bca68"
	actualRoot := root.Text(16)

	if actualRoot != expectedRoot {
		t.Fatalf("expected root %s, got %s", expectedRoot, actualRoot)
	}
}

func TestSetTxEffectivePercentage(t *testing.T) {
	smt := smt.NewSMT(nil)
	txIndex := big.NewInt(1)
	egp := common.HexToHash("0x000000000000000000000000000000005Ca1aB1E").Big()

	root, err := setTxEffectivePercentage(smt, txIndex, egp)
	if err != nil {
		t.Fatal(err)
	}

	// root taken from JS implementation
	expectedRoot := "f6b3130ecdd23bd9e47c4dda0fdde6bd0e0446c6d6927778e57e80016fa9fa23"
	actualRoot := root.Text(16)

	if actualRoot != expectedRoot {
		t.Fatalf("expected root %s, got %s", expectedRoot, actualRoot)
	}
}

func TestSetTxLogs(t *testing.T) {
	smt := smt.NewSMT(nil)
	txIndex := big.NewInt(1)
	logIndex := big.NewInt(1)
	log := common.HexToHash("0x000000000000000000000000000000005Ca1aB1E").Big()

	root, err := setTxLog(smt, txIndex, logIndex, log)
	if err != nil {
		t.Fatal(err)
	}

	// root taken from JS implementation
	expectedRoot := "aff38141ae4538baf61f08efe3019ef2d219f30b98b1d40a9813d502f6bacb12"
	actualRoot := root.Text(16)

	if actualRoot != expectedRoot {
		t.Fatalf("expected root %s, got %s", expectedRoot, actualRoot)
	}
}
