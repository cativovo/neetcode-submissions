class Solution {
    /**
     * @param {character[][]} board
     * @return {boolean}
     */
    isValidSudoku(board) {
       const squareMap = {};
    const colMap = {};
    const empty = ".";

    for (let i = 0; i < board.length; i++) {
      const rowMap = {};
      for (let j = 0; j < board[i].length; j++) {
        const value = board[i][j];
        if (value === empty) {
          continue;
        }

        // check for rows
        if (rowMap[value]) {
          return false;
        }

        rowMap[value] = true;

        // check for cols
        if (colMap[j] === undefined) {
          colMap[j] = {};
        }

        if (colMap[j][value]) {
          return false;
        }

        colMap[j][value] = true;

        // check for squares
        const squareId = Math.floor(i / 3) + Math.floor(j / 3) * 3;
        if (squareMap[squareId] === undefined) {
          squareMap[squareId] = {};
        }

        if (squareMap[squareId][value]) {
          return false;
        }

        squareMap[squareId][value] = true;
      }
    }

    return true;

    }
}
