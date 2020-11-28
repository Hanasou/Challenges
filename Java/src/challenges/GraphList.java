package challenges;

import java.util.Map;

public class GraphList implements Graph{
    class Node {
        private String id;
        private Map<Node, Integer> children;

        public String getId() {
            return this.id;
        }
        public Map<Node, Integer> getChildren() {
            return this.children;
        }
    }
    private Map<String, Node> nodes;

    @Override
    public void bfs() {

    }

    @Override
    public void dfs() {

    }
}
