#!/usr/bin/env python
import networkx as nx
import matplotlib.pyplot as plt
import numpy as np
import sys, getopt
from string import Template

def show_help():
    print('Usage:\n')
    print('    -f | --format    Set format. Default to "pdf" - Availables (pdf, png)')
    print('    -i | --input     Set input file (required)')
    print('    -o | --output    Set output file (required)')
    print('    -h               Show this help')

inputfile = ''
outputfile = ''
fileformat = 'pdf'
try:
    opts, args = getopt.getopt(sys.argv[1:], "hi:o:",["format=", "input=","output="])
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
            print('ERROR: Incorrect format - Availables (pdf, png)\n')
            show_help()
            sys.exit(3)
        fileformat = arg
if (outputfile == ''):
    print('ERROR: Missing input or output file\n')
    show_help()
    sys.exit(1)

if (inputfile == ''):
    infile = sys.stdin
else:
    infile = open(inputfile)

print("Reading data...")
data = []
for line in infile:
    if (line=='q'):
        break
    if (len(line) > 1):
        f1, f2 = line.split(",")
        data.append((f1.rstrip(), f2.rstrip()))
outputfile = outputfiletemplate.substitute(path=outputfile, format=fileformat)
g = nx.Graph()


print("Making graph...")
g.add_edges_from(data)
pos = nx.spring_layout(g, k=0.3*1/np.sqrt(len(g.nodes())), iterations=1)
plt.figure(3, figsize=(30, 30))
nx.draw(g, pos=pos)
nx.draw_networkx_labels(g, pos=pos)

print("Saving at", outputfile)
plt.savefig(outputfile)
