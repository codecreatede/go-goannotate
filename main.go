package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

/*
Author Gaurav Sablok
Universitat Potsdam
Date 2024-10-30

golang annotate and summarize your genome from the gtf or the gff file.

*/

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	os.Exit(1)
}

var annotationfile string

var rootCmd = &cobra.Command{
	Use:  "golanannotate",
	Long: "annotate and visualize your genome",
	Run:  annotateFunc,
}

func init() {
	rootCmd.Flags().
		StringVarP(&annotationfile, "annotationfile", "A", "path to the annotation file", "genome annotation")
}

func annotateFunc(cmd *cobra.Command, args []string) {
	type mRNADetails struct {
		mRNAParent string
		mRNAStrand string
		mRNAStart  int
		mRNAEnd    int
	}

	type exonDetails struct {
		exonParent string
		exonStrand string
		exonStart  int
		exonEnd    int
	}

	type proteinDetails struct {
		proteinParent string
		proteinStrand string
		proteinStart  int
		proteinEnd    int
	}

	type cdsDetails struct {
		cdsParent string
		cdsStrand string
		cdsStart  int
		cdsEnd    int
	}

	type fiveDetails struct {
		fiveParent string
		fiveStrand string
		fiveStart  int
		fiveEnd    int
	}

	type threeDetails struct {
		threeParent string
		threeStrand string
		threeStart  int
		threeEnd    int
	}

	mRNADet := []mRNADetails{}
	exonDet := []exonDetails{}
	cdsDet := []cdsDetails{}
	proteinDet := []proteinDetails{}
	threeDet := []threeDetails{}
	fiveDet := []fiveDetails{}

	annotateOpen, err := os.Open(annotationfile)
	if err != nil {
		log.Fatal(err)
	}
	annotateRead := bufio.NewScanner(annotateOpen)

	for annotateRead.Scan() {
		line := annotateRead.Text()
		if strings.Split(line, "\t")[2] == "mRNA" {
			start, _ := strconv.Atoi(strings.Split(string(line), "\t")[3])
			end, _ := strconv.Atoi(strings.Split(string(line), "\t")[4])
			mRNADet = append(mRNADet, mRNADetails{
				mRNAParent: strings.Split(string(line), "\t")[2],
				mRNAStrand: strings.Split(string(line), "\t")[6],
				mRNAStart:  start,
				mRNAEnd:    end,
			})
		}
		if strings.Split(line, "\t")[2] == "exon" {
			start, _ := strconv.Atoi(strings.Split(string(line), "\t")[3])
			end, _ := strconv.Atoi(strings.Split(string(line), "\t")[4])
			exonDet = append(exonDet, exonDetails{
				exonParent: strings.Split(string(line), "\t")[2],
				exonStrand: strings.Split(string(line), "\t")[6],
				exonStart:  start,
				exonEnd:    end,
			})
		}
		if strings.Split(line, "\t")[2] == "CDS" {
			start, _ := strconv.Atoi(strings.Split(string(line), "\t")[3])
			end, _ := strconv.Atoi(strings.Split(string(line), "\t")[4])
			cdsDet = append(cdsDet, cdsDetails{
				cdsParent: strings.Split(string(line), "\t")[2],
				cdsStrand: strings.Split(string(line), "\t")[6],
				cdsStart:  start,
				cdsEnd:    end,
			})
		}
		if strings.Split(line, "\t")[2] == "protein" {
			start, _ := strconv.Atoi(strings.Split(string(line), "\t")[3])
			end, _ := strconv.Atoi(strings.Split(string(line), "\t")[4])
			proteinDet = append(proteinDet, proteinDetails{
				proteinParent: strings.Split(string(line), "\t")[2],
				proteinStrand: strings.Split(string(line), "\t")[6],
				proteinStart:  start,
				proteinEnd:    end,
			})
		}
		if strings.Split(line, "\t")[2] == "five_prime_UTR" {
			start, _ := strconv.Atoi(strings.Split(string(line), "\t")[3])
			end, _ := strconv.Atoi(strings.Split(string(line), "\t")[4])
			fiveDet = append(fiveDet, fiveDetails{
				fiveParent: strings.Split(string(line), "\t")[2],
				fiveStrand: strings.Split(string(line), "\t")[6],
				fiveStart:  start,
				fiveEnd:    end,
			})
		}

		if strings.Split(line, "\t")[2] == "three_prime_UTR" {
			start, _ := strconv.Atoi(strings.Split(string(line), "\t")[3])
			end, _ := strconv.Atoi(strings.Split(string(line), "\t")[4])
			threeDet = append(threeDet, threeDetails{
				threeParent: strings.Split(string(line), "\t")[2],
				threeStrand: strings.Split(string(line), "\t")[6],
				threeStart:  start,
				threeEnd:    end,
			})
		}
	}

	exonLengthPlot := []int{}
	mRNALengthPlot := []int{}
	cdsLengthPlot := []int{}
	proteinLengthPlot := []int{}
	threeLengthPlot := []int{}
	fiveLengthPlot := []int{}

	for i := range exonDet {
		exonLengthPlot = append(exonLengthPlot, exonDet[i].exonEnd-exonDet[i].exonStart)
	}

	for i := range mRNADet {
		mRNALengthPlot = append(mRNALengthPlot, mRNADet[i].mRNAEnd-mRNADet[i].mRNAStart)
	}

	for i := range cdsDet {
		cdsLengthPlot = append(cdsLengthPlot, cdsDet[i].cdsEnd-cdsDet[i].cdsStart)
	}

	for i := range proteinDet {
		proteinLengthPlot = append(
			proteinLengthPlot,
			proteinDet[i].proteinEnd-proteinDet[i].proteinStart,
		)
	}

	for i := range threeDet {
		threeLengthPlot = append(threeLengthPlot, threeDet[i].threeEnd-threeDet[i].threeStart)
	}

	for i := range fiveDet {
		fiveLengthPlot = append(fiveLengthPlot, fiveDet[i].fiveEnd-fiveDet[i].fiveStart)
	}

	exonPlusLengthPlot := []int{}
	mRNAPlusLengthPlot := []int{}
	cdsPlusLengthPlot := []int{}
	proteinPlusLengthPlot := []int{}
	threePlusLengthPlot := []int{}
	fivePlusLengthPlot := []int{}

	exonMinusLengthPlot := []int{}
	mRNAMinusLengthPlot := []int{}
	cdsMinusLengthPlot := []int{}
	proteinMinusLengthPlot := []int{}
	threeMinusLengthPlot := []int{}
	fiveMinusLengthPlot := []int{}

	for i := range exonDet {
		if exonDet[i].exonStrand == "+" {
			exonPlusLengthPlot = append(exonPlusLengthPlot, exonDet[i].exonEnd-exonDet[i].exonStart)
		}
		if exonDet[i].exonStrand == "-" {
			exonMinusLengthPlot = append(
				exonMinusLengthPlot,
				exonDet[i].exonEnd-exonDet[i].exonStart,
			)
		}
	}

	for i := range mRNADet {
		if mRNADet[i].mRNAStrand == "+" {
			mRNAPlusLengthPlot = append(mRNAPlusLengthPlot, mRNADet[i].mRNAEnd-mRNADet[i].mRNAStart)
		}
		if mRNADet[i].mRNAStrand == "-" {
			mRNAMinusLengthPlot = append(
				mRNAMinusLengthPlot,
				mRNADet[i].mRNAEnd-mRNADet[i].mRNAStart,
			)
		}
	}

	for i := range cdsDet {
		if cdsDet[i].cdsStrand == "+" {
			cdsPlusLengthPlot = append(cdsPlusLengthPlot, cdsDet[i].cdsEnd-cdsDet[i].cdsStart)
		}
		if cdsDet[i].cdsStrand == "-" {
			cdsMinusLengthPlot = append(
				cdsMinusLengthPlot,
				cdsDet[i].cdsEnd-cdsDet[i].cdsStart,
			)
		}
	}

	for i := range fiveDet {
		if fiveDet[i].fiveStrand == "+" {
			fivePlusLengthPlot = append(fivePlusLengthPlot, fiveDet[i].fiveEnd-fiveDet[i].fiveStart)
		}
		if fiveDet[i].fiveStrand == "-" {
			fiveMinusLengthPlot = append(
				fiveMinusLengthPlot,
				fiveDet[i].fiveEnd-fiveDet[i].fiveStart,
			)
		}
	}

	for i := range threeDet {
		if threeDet[i].threeStrand == "+" {
			threePlusLengthPlot = append(
				threePlusLengthPlot,
				threeDet[i].threeEnd-threeDet[i].threeStart,
			)
		}
		if threeDet[i].threeStrand == "-" {
			threeMinusLengthPlot = append(
				threeMinusLengthPlot,
				threeDet[i].threeEnd-threeDet[i].threeStart,
			)
		}
	}

	for i := range proteinDet {
		if proteinDet[i].proteinStrand == "+" {
			proteinPlusLengthPlot = append(
				proteinPlusLengthPlot,
				proteinDet[i].proteinEnd-proteinDet[i].proteinStart,
			)
		}
		if proteinDet[i].proteinStrand == "-" {
			proteinMinusLengthPlot = append(
				proteinMinusLengthPlot,
				proteinDet[i].proteinEnd-proteinDet[i].proteinStart,
			)
		}
	}

	mRNASum := 0
	mRNAPlusSum := 0
	mRNAMinusSum := 0
	exonSum := 0
	exonPlusSum := 0
	exonMinusSum := 0
	proteinSum := 0
	proteinPlusSum := 0
	proteinMinusSum := 0
	cdsSum := 0
	cdsPlusSum := 0
	cdsMinusSum := 0
	fiveSum := 0
	fivePlusSum := 0
	fiveMinusSum := 0
	threeSum := 0
	threePlusSum := 0
	threeMinusSum := 0

	for i := 0; i <= len(mRNALengthPlot)-1; i++ {
		mRNASum += mRNALengthPlot[i]
	}
	for i := 0; i <= len(mRNAPlusLengthPlot)-1; i++ {
		mRNAPlusSum += mRNAPlusLengthPlot[i]
	}
	for i := 0; i <= len(mRNAMinusLengthPlot)-1; i++ {
		mRNAMinusSum += mRNAMinusLengthPlot[i]
	}
	for i := 0; i <= len(exonLengthPlot)-1; i++ {
		exonSum += exonLengthPlot[i]
	}
	for i := 0; i <= len(exonPlusLengthPlot)-1; i++ {
		exonPlusSum += exonPlusLengthPlot[i]
	}
	for i := 0; i <= len(exonMinusLengthPlot)-1; i++ {
		exonMinusSum += exonMinusLengthPlot[i]
	}

	for i := 0; i <= len(cdsLengthPlot)-1; i++ {
		cdsSum += cdsLengthPlot[i]
	}
	for i := 0; i <= len(proteinPlusLengthPlot)-1; i++ {
		cdsPlusSum += cdsPlusLengthPlot[i]
	}
	for i := 0; i <= len(proteinMinusLengthPlot)-1; i++ {
		cdsMinusSum += cdsMinusLengthPlot[i]
	}

	for i := 0; i <= len(proteinLengthPlot)-1; i++ {
		proteinSum += proteinLengthPlot[i]
	}
	for i := 0; i <= len(proteinPlusLengthPlot)-1; i++ {
		proteinPlusSum += proteinPlusLengthPlot[i]
	}
	for i := 0; i <= len(proteinMinusLengthPlot)-1; i++ {
		proteinMinusSum += proteinMinusLengthPlot[i]
	}

	for i := 0; i <= len(fiveLengthPlot)-1; i++ {
		fiveSum += fiveLengthPlot[i]
	}
	for i := 0; i <= len(fivePlusLengthPlot)-1; i++ {
		fivePlusSum += fivePlusLengthPlot[i]
	}
	for i := 0; i <= len(fiveMinusLengthPlot)-1; i++ {
		fiveMinusSum += fiveMinusLengthPlot[i]
	}

	for i := 0; i <= len(threeLengthPlot)-1; i++ {
		threeSum += threeLengthPlot[i]
	}
	for i := 0; i <= len(threePlusLengthPlot)-1; i++ {
		threePlusSum += threePlusLengthPlot[i]
	}
	for i := 0; i <= len(threeMinusLengthPlot)-1; i++ {
		threeMinusSum += threeMinusLengthPlot[i]
	}

	fmt.Println("The assembled genome from the gff annotations summare are given below:")
	fmt.Println("The total assembled mRNA are:", mRNASum)
	fmt.Println("The total assembled mRNA plus strand are:", mRNAPlusSum)
	fmt.Println("The total assembled mRNA minus strand are:", mRNAMinusSum)

	fmt.Println("The total assembled exon are:", exonSum)
	fmt.Println("The total assembled exon plus strand are:", exonPlusSum)
	fmt.Println("The total assembled exon minus strand are:", exonMinusSum)
	fmt.Println("The total assembled cds are:", cdsSum)
	fmt.Println("The total assembled cds plus strand are:", cdsPlusSum)
	fmt.Println("The total assembled cds minus strand are:", cdsMinusSum)
	fmt.Println("The total assembled protein are:", proteinSum)
	fmt.Println("The total assembled protein plus strand are:", proteinPlusSum)
	fmt.Println("The total assembled protein minus strand are:", proteinMinusSum)
	fmt.Println("The total assembled five are:", fiveSum)
	fmt.Println("The total assembled five plus strand are:", fivePlusSum)
	fmt.Println("The total assembled five minus strand are:", fiveMinusSum)
	fmt.Println("The total assembled three are:", threeSum)
	fmt.Println("The total assembled three plus strand are:", threePlusSum)
	fmt.Println("The total assembled three minus strand are:", threeMinusSum)

	mRNAfile, err := os.Create("genome-mRNA-stats.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer mRNAfile.Close()
	mRNAfile.WriteString("The stats on the mRNA are as follows" + "\n")
	mRNAfile.WriteString("LengthEstimates on the mRNA" + "\n")
	for i := range mRNALengthPlot {
		mRNAfile.WriteString(strconv.Itoa(mRNALengthPlot[i]) + "\n")
	}
	mRNAfile.WriteString("Length estimates on the plus strand are:" + "\n")
	for i := range mRNAPlusLengthPlot {
		mRNAfile.WriteString(strconv.Itoa(mRNAPlusLengthPlot[i]) + "\n")
	}
	mRNAfile.WriteString("Length estimates on the negative strand" + "\n")
	for i := range mRNAMinusLengthPlot {
		mRNAfile.WriteString(strconv.Itoa(mRNAMinusLengthPlot[i]) + "\n")
	}

	cdsfile, err := os.Create("genome-cds-stats.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer cdsfile.Close()
	cdsfile.WriteString("The stats on the cds are as follows" + "\n")
	cdsfile.WriteString("LengthEstimates on the cds" + "\n")
	for i := range cdsLengthPlot {
		cdsfile.WriteString(strconv.Itoa(cdsLengthPlot[i]) + "\n")
	}
	cdsfile.WriteString("Length estimates on the plus strand are:" + "\n")
	for i := range cdsPlusLengthPlot {
		cdsfile.WriteString(strconv.Itoa(cdsPlusLengthPlot[i]) + "\n")
	}
	cdsfile.WriteString("Length estimates on the negative strand" + "\n")
	for i := range cdsMinusLengthPlot {
		cdsfile.WriteString(strconv.Itoa(cdsMinusLengthPlot[i]) + "\n")
	}

	proteinfile, err := os.Create("genome-protein-stats.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer proteinfile.Close()
	proteinfile.WriteString("The stats on the protein are as follows" + "\n")
	proteinfile.WriteString("LengthEstimates on the mRNA" + "\n")
	for i := range proteinLengthPlot {
		proteinfile.WriteString(strconv.Itoa(proteinLengthPlot[i]) + "\n")
	}
	proteinfile.WriteString("Length estimates on the plus strand are:" + "\n")
	for i := range proteinPlusLengthPlot {
		proteinfile.WriteString(strconv.Itoa(proteinPlusLengthPlot[i]) + "\n")
	}
	proteinfile.WriteString("Length estimates on the negative strand" + "\n")
	for i := range proteinMinusLengthPlot {
		proteinfile.WriteString(strconv.Itoa(proteinMinusLengthPlot[i]) + "\n")
	}

	exonfile, err := os.Create("genome-exon-stats.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer exonfile.Close()
	exonfile.WriteString("The stats on the exon are as follows" + "\n")
	exonfile.WriteString("LengthEstimates on the exon" + "\n")
	for i := range exonLengthPlot {
		exonfile.WriteString(strconv.Itoa(exonLengthPlot[i]) + "\n")
	}
	exonfile.WriteString("Length estimates on the plus strand are:" + "\n")
	for i := range exonPlusLengthPlot {
		exonfile.WriteString(strconv.Itoa(exonPlusLengthPlot[i]) + "\n")
	}
	exonfile.WriteString("Length estimates on the negative strand" + "\n")
	for i := range exonMinusLengthPlot {
		exonfile.WriteString(strconv.Itoa(exonMinusLengthPlot[i]) + "\n")
	}

	fivefile, err := os.Create("genome-five-stats.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fivefile.Close()
	fivefile.WriteString("The stats on the five prime UTR are as follows" + "\n")
	fivefile.WriteString("LengthEstimates on the five prime UTR" + "\n")
	for i := range fiveLengthPlot {
		fivefile.WriteString(strconv.Itoa(fiveLengthPlot[i]) + "\n")
	}
	fivefile.WriteString("Length estimates on the plus strand are:" + "\n")
	for i := range fivePlusLengthPlot {
		fivefile.WriteString(strconv.Itoa(fivePlusLengthPlot[i]) + "\n")
	}
	fivefile.WriteString("Length estimates on the negative strand" + "\n")
	for i := range fiveMinusLengthPlot {
		fivefile.WriteString(strconv.Itoa(fiveMinusLengthPlot[i]) + "\n")
	}

	threefile, err := os.Create("genome-three-stats.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer threefile.Close()
	threefile.WriteString("The stats on the three prime UTR are as follows" + "\n")
	threefile.WriteString("LengthEstimates on the mRNA" + "\n")
	for i := range mRNALengthPlot {
		threefile.WriteString(strconv.Itoa(threeLengthPlot[i]) + "\n")
	}
	threefile.WriteString("Length estimates on the plus strand are:" + "\n")
	for i := range threePlusLengthPlot {
		threefile.WriteString(strconv.Itoa(threePlusLengthPlot[i]) + "\n")
	}
	threefile.WriteString("Length estimates on the negative strand" + "\n")
	for i := range threeMinusLengthPlot {
		threefile.WriteString(strconv.Itoa(threeMinusLengthPlot[i]) + "\n")
	}
}
