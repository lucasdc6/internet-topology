#!/usr/bin/env python
import networkx as nx
import matplotlib.pyplot as plt
import numpy as np
import sys, getopt
from string import Template

def show_help():
    print 'Usage:\n'
    print '    -f | --format    Set format. Default to "pdf" - Availables (pdf, png)'
    print '    -i | --input     Set input file (required)'
    print '    -o | --output    Set output file (required)'
    print '    -h               Show this help'

def parse(data2):
    x = data2.split(",")
    if (len(x) == 2):
        return (x[0], x[1])

def main(argv):
    inputfile = ''
    outputfile = ''
    fileformat = 'pdf'
    try:
        opts, args = getopt.getopt(argv,"hi:o:",["format=", "input=","output="])
    except getopt.GetoptError:
        show_help()
        sys.exit(2)
    for opt, arg in opts:
        if opt == '-h':
            show_help()
            sys.exit()
        elif opt in ("-i", "--input"):
            inputfile = arg
        elif opt in ("-o", "--output"):
            outputfiletemplate = Template("$path.$format")
            outputfile = arg
        elif opt == '--format':
            if arg not in ("png", "pdf"):
                print 'ERROR: Incorrect format - Availables (pdf, png)\n'
                show_help()
                sys.exit(3)
            fileformat = arg
    if (inputfile == '' or outputfile == ''):
        print 'ERROR: Missing input or output file\n'
        show_help()
        sys.exit(1)
    outputfile = outputfiletemplate.substitute(path=outputfile, format=fileformat)
    data = open(inputfile).read().split("\n")
    asConnections = filter(lambda x: x is not None, map(parse, list(data)))
    g = nx.Graph()

    g.add_edges_from(asConnections)
    pos = nx.spring_layout(g, k=0.3*1/np.sqrt(len(g.nodes())), iterations=1)
    plt.figure(3, figsize=(30, 30))
    nx.draw(g, pos=pos)
    nx.draw_networkx_labels(g, pos=pos)
    plt.savefig(outputfile)

if __name__ == "__main__":
    main(sys.argv[1:])
