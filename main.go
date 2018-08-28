package main

import (
	"time"
	//"errors"
	//"context"
	"log"
	//"time"
	//"github.com/jlaffaye/ftp"
)

var files = []string{
	"http://adsfadasfafdsafasdfasf.com/",
	"http://www.python.org/ftp/python/2.7.3/Python-2.7.3.tgz",
	"http://sourceforge.net/projects/pycogent/files/PyCogent/1.5.3/PyCogent-1.5.3.tgz/download",
	"http://sourceforge.net/projects/numpy/files/NumPy/1.5.1/numpy-1.5.1.tar.gz/download",
	//"ftp://thebeast.colorado.edu/pub/biom-format-releases/biom-format-1.1.2.tar.gz",
	"http://www.drive5.com/uclust/uclustq1.2.22_i86linux64",
	// WAS "https://github.com/downloads/qiime/pynast/PyNAST-1.2.tar.gz",
	"https://github.com/biocore/pynast/archive/1.2.2.tar.gz",
	"http://greengenes.lbl.gov/Download/Sequence_Data/Fasta_data_files/core_set_aligned.fasta.imputed",
	"http://greengenes.lbl.gov/Download/Sequence_Data/lanemask_in_1s_and_0s",
	"http://www.microbesonline.org/fasttree/FastTree-2.1.3.c",
	//"http://java.sun.com/javase/downloads/index.jsp",
	"http://sourceforge.net/projects/rdp-classifier/files/rdp-classifier/rdp_classifier_2.2.zip/download",
	"https://downloads.sourceforge.net/project/tax2tree/tax2tree-v1.0.tar.gz",
	"http://mirrors.vbi.vt.edu/mirrors/ftp.ncbi.nih.gov/blast/executables/blast+/2.2.22/ncbi-blast-2.2.22%2B-src.tar.gz",
	"http://www.bioinformatics.org/download/cd-hit/cd-hit-2007-0131.tar.gz",
	"https://sourceforge.net/projects/microbiomeutil/files/__OLD_VERSIONS/microbiomeutil_2010-04-29.tar.gz/download",
	"http://www.mothur.org/w/images/6/6d/Mothur.1.25.0.zip",
	"http://www.mothur.org/w/images/9/91/Clearcut.source.zip",
	//"ftp://thebeast.colorado.edu/pub/QIIME-v1.5.0-dependencies/stamatak-standard-RAxML-5_7_2012.tgz",
	// WAS "ftp://selab.janelia.org/pub/software/infernal/infernal.tar.gz",
	"https://github.com/EddyRivasLab/infernal/archive/1.0.2.tar.gz",
	"ftp://occams.dfci.harvard.edu/pub/bio/tgi/software/cdbfasta/cdbfasta.tar.gz",
	"http://www.drive5.com/muscle/downloads.htm",
	"http://static.davidsoergel.com/rtax-0.983.tgz",
	"http://matsen.fhcrc.org/pplacer/builds/pplacer-v1.1-Linux.tar.gz",
	"http://downloads.sourceforge.net/project/parsinsert/ParsInsert.1.04.tgz",
	"ftp://ftp.gnu.org/gnu/gsl/gsl-1.9.tar.gz",
	// WAS "http://ampliconnoise.googlecode.com/files/AmpliconNoiseV1.27.tar.gz",
	"https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/ampliconnoise/AmpliconNoiseV1.27.tar.gz",
	"https://downloads.haskell.org/~ghc/8.4.3/ghc-8.4.3-src.tar.xz",
	"http://cran.utstat.utoronto.ca/src/base/R-2/R-2.12.0.tar.gz",
	// WAS "ftp://169.228.46.98/gg_12_10/gg_12_10_otus.tar.gz",
	"https://s3.amazonaws.com/gg_sg_web/gg_12_10_otus.tar.gz?AWSAccessKeyId=AKIAIKZRXPOMF7SLT42A&Signature=1m0L03Kfpi6XiB8I6NXKAHX5Ytw%3D&Expires=1534531563",
	"ftp://greengenes.microbio.me/greengenes_release/gg_12_10/gg_12_10_otus.tar.gzZZZZ",
	//"http://gdata-python-client.googlecode.com/files/gdata-2.0.17.tar.gz",
	"https://files.pythonhosted.org/packages/b2/e0/6e062327b211e9b1c5f30f65a9a65cf49eb1d3a7da3ce42fdc9a9e128535/gdata-2.0.17.tar.gz",
	// SHA256
}

const defaultNumFtpRoutines = 3
const defaultNumHttpRoutines = 6

func main() {
	var err error
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	urls, err := newUrls(files)
	if err != nil {
		log.Fatal(err)
	}
	//getHeadInfo(urls)

	for i, _ := range urls {
		l, err := urls[i].GetRemoteSize()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Size=", l, " ", urls[i])
		err = urls[i].SampleTime()
		if err != nil {
			log.Println(err)
			continue
		}
		if l > 0 && l < 1000000 {
			log.Println("DOWNLOADING " + urls[i].Url())
			r, err := urls[i].Get()
			if err != nil {
				log.Println(err)
				continue
			}
			writeFile(r, "foo", time.Now())
		}
	}

}
