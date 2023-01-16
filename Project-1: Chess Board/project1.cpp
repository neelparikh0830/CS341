#include "grader.h"
#include <string>
#include <iostream>
#include <fstream>

using namespace std;

class chessboard {
   private:
    struct Board {
        testcolor col = Car;
        testpiece pie = Zebra;  
    };

    Board nike[8][8];
    
    testcolor getColor(int x, int y) {
        testcolor r = nike[x][y].col;
        return r;
    }
    
    testpiece getPiece(int x, int y){
        testpiece e = nike[x][y].pie;
        return e;
    }
    
    public:
    int testPlace(int x, int y, testcolor c, testpiece p)
    {
        if (x >= 8 || x < 0)
        {
            return -1;
        }
        else if (y >= 8 || y < 0)
        {
            return -2;
        }
        else if (nike[x][y].col != Car)
        {
            return -3;
        }
        else if (c != White && c != Black)
        {
            return -4;
        }
        else if (p != Rook && p != Knight && p != Bishop && p != Queen && p != King && p != Pawn)
        {
            return -5;
        }
        nike[x][y].col = c;
        nike[x][y].pie = p;
        return 1;
    }
    
    int testGet(int x, int y, testcolor &c, testpiece &p){
        if (x >= 8 || x < 0)
        {
            return -1;
        }
        else if (y >= 8 || y < 0)
        {
            return -2;
        }
        else if (nike[x][y].pie == Zebra)
        {
            return -3;
        }
        c = getColor(x, y);
        p = getPiece(x, y);
        return 1;
    }
   
};