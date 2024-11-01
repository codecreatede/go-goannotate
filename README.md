# go-goannotate

- annotation and genome annotation summarizer.
- takes a gff file and summarizes all the regions such as mRNA, cds, protein five prime UTR, and three prime UTR. 
- process as many genomes at a single time 
```
git clone htts.github.com/go-goannotate
go run main.go

```
- detail usage 

```
╰─$ go run main.go -h
annotate and visualize your genome

Usage:
  golanannotate [flags]

Flags:
  -A, --annotationfile string   genome annotation (default "path to the annotation file")
  -h, --help                    help for golanannotate
exit status 1
```
- it will generate the summary 
```
╭─gauavsablok@gauravsablok ~/Desktop/go/go-goannotate ‹main●›
╰─$ go run main.go -A sample-file/sample-TAIR10_GFF3_genes.gff
The assembled genome from the gff annotations summare are given below:
The total assembled mRNA are: 26857
The total assembled mRNA plus strand are: 18053
The total assembled mRNA minus strand are: 8804
The total assembled exon are: 19536
The total assembled exon plus strand are: 13976
The total assembled exon minus strand are: 5560
The total assembled cds are: 15712
The total assembled cds plus strand are: 552
The total assembled cds minus strand are: 276
The total assembled protein are: 22456
The total assembled protein plus strand are: 16990
The total assembled protein minus strand are: 5466
The total assembled five are: 1477
The total assembled five plus strand are: 602
The total assembled five minus strand are: 875
The total assembled three are: 2128
The total assembled three plus strand are: 455
The total assembled three minus strand are: 1673
exit status 1

```
Gaurav Sablok
