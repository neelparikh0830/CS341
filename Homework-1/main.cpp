//
// Homework 1
// CS 341 Spring 2021
// << Neel Parikh>>
//

#include <iostream>
#include <fstream>
#include <istream>
using namespace std;

int main()
{
	string file1, file2, line;
	cin >> file1;
	ifstream myFile (file1);
	if (myFile.is_open()) 
	{
		while (getline (myFile,line)) 
		{
			cout<<line<<'\n';
		}
		ifstream readFile (line);
		if (readFile.is_open()) 
		{
			while (getline (readFile,file2)) 
			{
				cout<<file2<<'\n';
			}
		}
		readFile.close();
	}
	myFile.close();
	return 0;
}
