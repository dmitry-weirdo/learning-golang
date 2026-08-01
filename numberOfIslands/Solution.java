package numberOfIslands;

class Solution {
    // Instance variables to store grid and its dimensions
    private char[][] grid;
    private int rows;
    private int cols;

    /**
     * Counts the number of islands in a 2D grid.
     * An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
     * 
     * @param grid 2D grid map of '1's (land) and '0's (water)
     * @return the number of islands
     */
    public int numIslands(char[][] grid) {
        // Initialize grid dimensions
        this.rows = grid.length;
        this.cols = grid[0].length;
        this.grid = grid;
      
        int islandCount = 0;
      
        // Traverse each cell in the grid
        for (int row = 0; row < rows; row++) {
            for (int col = 0; col < cols; col++) {
                // If current cell is land ('1'), start DFS to mark entire island
                if (grid[row][col] == '1') {
                    dfs(row, col);
                    islandCount++;
                }
            }
        }
      
        return islandCount;
    }

    /**
     * Performs depth-first search to mark all connected land cells as visited.
     * Marks visited land cells by changing them from '1' to '0'.
     * 
     * @param row current row index
     * @param col current column index
     */
    private void dfs(int row, int col) {
        // Mark current land cell as visited by setting it to water ('0')
        grid[row][col] = '0';
      
        // Direction vectors for exploring 4 adjacent cells (up, right, down, left)
        // Using pairs: (-1,0), (0,1), (1,0), (0,-1)
        int[] directions = {-1, 0, 1, 0, -1};
      
        // Explore all 4 adjacent cells
        for (int i = 0; i < 4; i++) {
            int newRow = row + directions[i];
            int newCol = col + directions[i + 1];
          
            // Check if adjacent cell is within bounds and is unvisited land
            if (newRow >= 0 && newRow < rows && 
                newCol >= 0 && newCol < cols && 
                grid[newRow][newCol] == '1') {
                // Recursively explore the adjacent land cell
                dfs(newRow, newCol);
            }
        }
    }
}