#include<iostream>

/*
design an algorithm that determines the maximum profit that
could have been made by buying and then selling a single share over a given day
range, subject to the constraint that the buy and the sell have to take place at the start
of the day. (This algorithm may be needed to backtest a trading strategy.)
*/

/*
Iterate through S, keeping track of the minimum element m seen thus far. If the
difference of the current element and m is greater than the maximum profit recorded
so far, update the maximum profit. This algorithm performs a constant amount of
work per array element, leading to an O(n) time complexity.
*/

int main()
{
int input[] = {9, 4, 3, 6, 5, 8, 1, 4, 3, 2, 8, 9, 4};
int i = 0;
int currMinIndex = i;
int currMaxIndex = i;
int candidateMinIndex = i;
int currMin = input[i];
int currMax = input[i];
int candidateMin = input[i];
std::cout << sizeof(input);
int length = 13;
while(i < length-1){
    

        if ((input[i+1] - input[i]) > (currMax - currMin))
        {
            currMax = input[i+1];
            currMin = input[i];
            currMaxIndex = i+1;
            currMinIndex = i;
        }
        if (input[i] < candidateMin){
            candidateMin = input[i];
            candidateMinIndex = i;
        }
    
    if((input[i] - candidateMin) > (currMax - currMin))
    {
            currMax = input[i];
            currMin = candidateMin;
            currMaxIndex = i;
            currMinIndex = candidateMinIndex;
    }
    
    
    i++;
}

std::cout << currMaxIndex << currMinIndex;
}
