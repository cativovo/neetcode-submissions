class Solution {
    /**
     * @param {character[][]} board
     * @return {boolean}
     */
    isValidSudoku(board) {
            const squaresMap = {};
    const colMap = {};
    const empty = ".";

    for (let i = 0; i < board.length; i++) {
      const rowMap = {};
      for (let j = 0; j < board[i].length; j++) {
        const colValue = board[i][j];
        if (colValue === empty) {
          continue;
        }

        // check for rows
        if (rowMap[colValue]) {
          return false;
        }

        rowMap[colValue] = true;

        // check for cols
        if (colMap[j] === undefined) {
          colMap[j] = {};
        }

        if (colMap[j][colValue]) {
          return false;
        }

        colMap[j][colValue] = true;

        // check for squares
        const squareId = `${Math.floor(i / 3)}${Math.floor(j / 3)}`;
        if (squaresMap[squareId] === undefined) {
          squaresMap[squareId] = {};
        }

        if (squaresMap[squareId][colValue]) {
          return false;
        }

        squaresMap[squareId][colValue] = true;
      }
    }

    return true;

    }
}
