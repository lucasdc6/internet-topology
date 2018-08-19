import networkx as nx
import matplotlib.pyplot as plt
import numpy as np

def parse(data2):
    x = data2.split(",")
    if (len(x) == 2):
        return (x[0], x[1])

data = open("/tmp/internet-topology").read().split("\n")
asConnections = filter(lambda x: x is not None, map(parse, list(data)))
g = nx.Graph()

g.add_edges_from(asConnections)
pos = nx.spring_layout(g, k=0.3*1/np.sqrt(len(g.nodes())), iterations=1)
plt.figure(3, figsize=(30, 30))
nx.draw(g, pos=pos)
nx.draw_networkx_labels(g, pos=pos)
plt.savefig("path_full.pdf")
