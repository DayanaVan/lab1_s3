#pragma once

#include <iostream>
#include <fstream>
#include <algorithm>
#include "structures.h"

#include <iostream>

int readArray(std::ifstream &ifs, Array *arr);

void writeArray(std::ofstream &ofs, Array *arr);

int readForwardList(std::ifstream &ifs, ForwardList *fl);

void writeForwardList(std::ofstream &ofs, ForwardList *fl);

int readList(std::ifstream &ifs, List *l);

void writeList(std::ofstream &ofs, List *l);

int readStack(std::ifstream &ifs, Stack *s);

void writeStackRec(std::ofstream &ofs, StackNode *sn);

void writeStack(std::ofstream &ofs, Stack *s);

int readQueue(std::ifstream &ifs, Queue *q);

void writeQueue(std::ofstream &ofs, Queue *q);

int readTreeNode(std::ifstream &ifs, TreeNode *&node);

int readTree(std::ifstream &ifs, Tree *t);

void writeTree(std::ofstream &ofs, Tree *t);

bool isNumber(std::string s);
