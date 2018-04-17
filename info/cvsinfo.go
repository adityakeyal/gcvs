package cvsinfo

/*
* CVSInfo... is a struct holding the Information required for CVS
 */
type CVSInfo struct {
	CVSRoot string

	NonRecursiveCommand []string
	RecursiveCommand    []string
	NonRecursiveModule  []string
	RecursiveModule     []string

	// // -command for
	// var mycommandonlydir = []string{"-d", CVSRoot, "checkout", "-P", "-l", "-r", tagName, "--"}
	// var mycommandrecur = []string{"-d", CVSRoot, "checkout", "-P", "-r", tagName, "--"}

	// var foldersonly = []string{"istar-gv/devel/", "istar-gv/devel/gv-core/", "istar-gv/devel/gv-deployments/"}

	// var recursive = []string{"istar-gv/devel/gv-core/gv-core-adp", "istar-gv/devel/gv-core/gv-core-bkg", "istar-gv/devel/gv-core/gv-core-cam", "istar-gv/devel/gv-core/gv-core-cax", "istar-gv/devel/gv-core/gv-core-cli", "istar-gv/devel/gv-core/gv-core-cnf", "istar-gv/devel/gv-core/gv-core-dbtest", "istar-gv/devel/gv-core/gv-core-dist", "istar-gv/devel/gv-core/gv-core-dro", "istar-gv/devel/gv-core/gv-core-drv", "istar-gv/devel/gv-core/gv-core-exm", "istar-gv/devel/gv-core/gv-core-frx", "istar-gv/devel/gv-core/gv-core-gle", "istar-gv/devel/gv-core/gv-core-gwy", "istar-gv/devel/gv-core/gv-core-inf", "istar-gv/devel/gv-core/gv-core-itf", "istar-gv/devel/gv-core/gv-core-mtf", "istar-gv/devel/gv-core/gv-core-ncm", "istar-gv/devel/gv-core/gv-core-orb", "istar-gv/devel/gv-core/gv-core-pmt", "istar-gv/devel/gv-core/gv-core-rec", "istar-gv/devel/gv-core/gv-core-ref", "istar-gv/devel/gv-core/gv-core-rst", "istar-gv/devel/gv-core/gv-core-rttm", "istar-gv/devel/gv-core/gv-core-slr", "istar-gv/devel/gv-core/gv-core-srs", "istar-gv/devel/gv-core/gv-core-stl", "istar-gv/devel/gv-core/gv-core-tax", "istar-gv/devel/gv-core/gv-core-test", "istar-gv/devel/gv-core/gv-core-trd", "istar-gv/devel/gv-deployments/gv-dist", "istar-gv/devel/gv-deployments/gv-th"}
}

/**
Constructor :: NewCVSInfo
*/
func NewCVSInfo(tagName string) CVSInfo {
	cvsInfo := CVSInfo{}

	cvsInfo.CVSRoot = ":ssh;username=adityak;hostname=cvs.nrifintech.com;privatekey='d:\\cvs-config\\newcvsprivatekey.ppk':/home/secure/services/cvsroot/istar-gv/istar-gvrepo"
	cvsInfo.NonRecursiveCommand = []string{"-d", cvsInfo.CVSRoot, "checkout", "-P", "-l", "-r", tagName, "--"}
	cvsInfo.RecursiveCommand = []string{"-d", cvsInfo.CVSRoot, "checkout", "-P", "-r", tagName, "--"}

	cvsInfo.NonRecursiveModule = []string{"istar-gv/devel/", "istar-gv/devel/gv-core/", "istar-gv/devel/gv-deployments/"}
	cvsInfo.RecursiveModule = []string{"istar-gv/devel/gv-core/gv-core-adp",
		"istar-gv/devel/gv-core/gv-core-bkg", "istar-gv/devel/gv-core/gv-core-cam", "istar-gv/devel/gv-core/gv-core-cax",
		"istar-gv/devel/gv-core/gv-core-cli", "istar-gv/devel/gv-core/gv-core-cnf", "istar-gv/devel/gv-core/gv-core-dbtest",
		"istar-gv/devel/gv-core/gv-core-dist", "istar-gv/devel/gv-core/gv-core-dro", "istar-gv/devel/gv-core/gv-core-drv",
		"istar-gv/devel/gv-core/gv-core-exm", "istar-gv/devel/gv-core/gv-core-frx", "istar-gv/devel/gv-core/gv-core-gle",
		"istar-gv/devel/gv-core/gv-core-gwy", "istar-gv/devel/gv-core/gv-core-inf", "istar-gv/devel/gv-core/gv-core-itf",
		"istar-gv/devel/gv-core/gv-core-mtf", "istar-gv/devel/gv-core/gv-core-ncm", "istar-gv/devel/gv-core/gv-core-orb",
		"istar-gv/devel/gv-core/gv-core-pmt", "istar-gv/devel/gv-core/gv-core-rec", "istar-gv/devel/gv-core/gv-core-ref",
		"istar-gv/devel/gv-core/gv-core-rst", "istar-gv/devel/gv-core/gv-core-rttm", "istar-gv/devel/gv-core/gv-core-slr",
		"istar-gv/devel/gv-core/gv-core-srs", "istar-gv/devel/gv-core/gv-core-stl", "istar-gv/devel/gv-core/gv-core-tax",
		"istar-gv/devel/gv-core/gv-core-test", "istar-gv/devel/gv-core/gv-core-trd", "istar-gv/devel/gv-deployments/gv-dist",
		"istar-gv/devel/gv-deployments/gv-th"}

	return cvsInfo

}
