package main

import "flag"

/*
Author Gaurav Sablok
Universitat Potsdam
Date 2024-9-12

golang for the graphics to visualize the gene structure and the genome structure coming from the
genome annotations such as plant, bacterial and fungal genome anntoations. Check the sample genome annotations
as how it should look for the go-graphic to draw the genome annotation.

*/

func main() {
	annotateGFF := flag.String("GFFannotation", "enter the path to the GFF annotation file", "file")
	annotateGTF := flag.String("GTFannotation", "enter the path to the GTF annotation file", "file")

	flag.Parse()
}
