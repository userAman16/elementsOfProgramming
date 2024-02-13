#include<iostream>

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