package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

/*
Author Gaurav Sablok
Universitat Potsdam
Date 2024-9-12

golang for the graphics to visualize the gene structure and the genome structure coming from the
genome annotations such as plant, bacterial and fungal genome anntoations. Check the sample genome annotations
as how it should look for the go-graphic to draw the genome annotation.

reads a genome annotation file, see the sample file and then stores all the informaiton in the structs and
then makes all the plots for the genome visualization for each genes and parent genes and also show the arc for the
over-lapping genes.

*/

func main() {
	annotateGFF := flag.String("GFFannotation", "enter the path to the GFF annotation file", "file")

	flag.Parse()

  type annotateGFFstruct struct {
     geneID string

	}

	type mRNADetails struct {
		mRNAParent string
		mRNAstart string
		mRNAend string
	}

	type exonDetails struct {
		exonParent string
		exonStart string
		exonEnd string
	}

	type proteinDetails struct {
		proteinParent string
		proteinStart string
		proteinEnd string
	}

	type cdsDetails struct {
     cdsParent string
		 cdsStart string
		 cdsEnd string
	}

	type fiveDetails struct {
		fiveParent string
		fiveStart string
		fiveEnd string
	}

	type threeDetails struct {
		threeParent string
		threeStart string
		threeEnd string
	}

	annoateOpen, err := os.Open(annotateGFF)
	if err != nil {
		log.Fatal(err)
	}


	annotateRead := bufio.NewScanner(annotateOpen)

	annotateID := []annotateGFFstruct{}
	for annotateRead.Scan() {
		line := annotateRead.Text()
		annoatateID = append(annotateID, annotateGFFstruct{
			geneID : strings.Split(string(strings.Split(strings.HasPrefix(string(line),"Parent"), ",")[0]), "=")[1],
		})
	}

    mRNADet := []mRNADetails{}
	for annotateRead.Scan() {
		for j := range annotateID {
		line := annotateRead.Text()
		id := strings.Split(string(strings.Split(strings.HasPrefix(string(line), "Parent"), ",")[0]), "=")[1]
		if annotateID[i].geneID == id && strings.Split(line, "\t")[2] == "mRNA" {
		mRNADet = append(mRNADet, mRNADetails{
			mRNAParent : strings.Split(string(strings.Split(strings.HasPrefix(string(line),"Parent"), ",")[0]), "=")[1],
			mRNAStart : strings.Split(string(line), "\t")[4],
			mRNAEnd : strings.Split(string(line), "\t")[5],
		})
	}
     }
  } 

    exonDet := []exonDetails{}
	for annotateRead.Scan() {
		for j := range annotateID {
		line := annotateRead.Text()
		id := strings.Split(string(strings.Split(strings.HasPrefix(string(line), "Parent"), ",")[0]), "=")[1]
		if annotateID[i].geneID == id && strings.Split(line, "\t")[2] == "exon" {
		exonDet = append(exonDet, exonDetails{
			exonParent : strings.Split(string(strings.Split(strings.HasPrefix(string(line),"Parent"), ",")[0]), "=")[1],
			exonStart : strings.Split(string(line),"\t")[4],
			exonEnd : strings.Split(string(line),"\t")[5],
		})
	}
}
	}

    cdsDet := []cdsDetails{}
	for annotateRead.Scan() {
		for j := range annotateID {
		line := annotateRead.Text()
		id := strings.Split(string(strings.Split(strings.HasPrefix(string(line), "Parent"), ",")[0]), "=")[1]
		if annotateID[i].geneID == id && strings.Split(line, "\t")[2] == "CDS" {
		cdsDet = append(cdsDet, cdsDetails{
			cdsParent : strings.Split(string(strings.Split(strings.HasPrefix(string(line),"Parent"), ",")[0]), "=")[1],
			cdsStart : strings.Split(string(line), "\t")[4],
			cdsEnd : strings.Split(string(line), "\t")[5],
		})
		}
    }
}
    proteinDet := []proteinDetails{}
	for annotateRead.Scan() {
		for j := range annotateID {
		line := annotateRead.Text()
		id := strings.Split(string(strings.Split(strings.HasPrefix(string(line), "Parent"), ",")[0]), "=")[1]
		if annotateID[i].geneID == id && strings.Split(line, "\t")[2] == "protein" {
		proteinDet = append(mRNADet, proteinDetails{
			proteinParent : strings.Split(string(strings.Split(strings.HasPrefix(string(line),"Parent"), ",")[0]), "=")[1],
			proteinStart : strings.Split(string(line), "\t")[4],
			proteinEnd : strings.Split(string(line), "\t")[5],
		})
		}
}
	}

   		fiveDet := []fiveDetails{}
   		for annotateRead.Scan() {
		for j := range annotateID {
		line := annotateRead.Text()
		id := strings.Split(string(strings.Split(strings.HasPrefix(string(line), "Parent"), ",")[0]), "=")[1]
		if annotateID[i].geneID == id && strings.Split(line, "\t")[2] == "five_prime_UTR" {
		fiveDet = append(mRNADet, fiveDetails{
			fiveParent : strings.Split(string(strings.Split(strings.HasPrefix(string(line),"Parent"), ",")[0]), "=")[1],
			fiveStart : strings.Split(string(line), "\t")[4],
			fiveEnd : strings.Split(string(line), "\t")[5],
		})
	}
}
   }

		threeDet := []threeDetails{}
        for annotateRead.Scan() {
		for j := range annotateID {
		line := annotateRead.Text()
		id := strings.Split(string(strings.Split(strings.HasPrefix(string(line), "Parent"), ",")[0]), "=")[1]
		if annotateID[i].geneID == id && strings.Split(line, "\t")[2] == "three_prime_UTR" {
		threeDet := []threeDetails{}
		threeDet = append(threeDet, threeDetails{
			threeParent : strings.Split(string(strings.Split(strings.HasPrefix(string(line),"Parent"), ",")[0]), "=")[1],
			threeStart : strings.Split(string(line), "\t)[4],
			threeEnd : strings.Split(string(line), "\t)[5],
		})
	}
}

// plotting features to add using the go graphics package for tomorrow

 

}


