package config

const (
	PojectRoot = "/home/dev/projects/dl-base"

	StaticFolderName    = "static"
	AssetFolderName     = StaticFolderName + "/assets"
	Filename            = "samples.json"
	SampleFilePath      = PojectRoot + "/" + AssetFolderName + "/" + Filename // /home/u/static/assets/samples.json
	IterScalar          = 500
	RecordeStripeScalar = 10

	IterVector          = 1000
	RecordeStripeVector = 10
)
