package config

const (
	PojectRoot = "/home/dev/projects/dl-base"

	StaticFolderName = "docs"
	AssetFolderName  = StaticFolderName + "/assets"

	csv3dFileName = "3d-scatter-gen.csv"
	CSV3dFile     = PojectRoot + "/" + StaticFolderName + "/dataset/" + csv3dFileName // /home/u/docs/dataset/3d-scatter-gen.csv

	JsonFilename   = "samples.json"
	SampleFilePath = PojectRoot + "/" + StaticFolderName + "/dataset/" + JsonFilename // /home/u/docs/samples.json

	IterScalar           = 500
	RecordeStripe4Scalar = 10

	IterVector           = 1000
	RecordeStripe4Vector = 100
)
