#include <stdio.h>
#include "InfInt.h"

InfInt factorial (int x) 
{
    InfInt nike = 1;
    while (x > 1)
    {
        nike *= x--;
    }
    return nike;
}

int main () 
{
    InfInt love = factorial (100);
    return 0;
}