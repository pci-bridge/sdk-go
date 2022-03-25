package pcib

type Permission int64

const (
	None                     Permission = 0
	CreateServiceAccount     Permission = 8
	UpdateServiceAccount     Permission = 16
	DeactivateServiceAccount Permission = 32
	RetrieveServiceAccount   Permission = 64
	ListServiceAccounts      Permission = 128
	CreatePlacement          Permission = 256
	UpdatePlacement          Permission = 512
	PausePlacement           Permission = 1024
	ResumePlacement          Permission = 2048
	DeactivatePlacement      Permission = 4096
	RetrievePlacement        Permission = 8192
	ListPlacements           Permission = 16384
	CreateEndpoint           Permission = 32768
	DeactivateEndpoint       Permission = 65536
	PauseEndpoint            Permission = 131072
	ResumeEndpoint           Permission = 262144
	RetrieveEndpoint         Permission = 524288
	ListEndpoints            Permission = 1048576
	Tokenize                 Permission = 2097152
	DeactivateToken          Permission = 4194304
	ReTokenize               Permission = 8388608
	ProxyRequest             Permission = 16777216
	RefreshRequest           Permission = 33554432
)
