#pragma once

#include <string>

struct List;

struct ListNode
{
    ListNode *next;
    ListNode *prev;
    std::string key;

    ListNode(std::string key);
};

struct List
{
    ListNode *tail;
    ListNode *head;

    List();

    void addHead(std::string key);

    void addTail(std::string key);

    void insert(std::string key, int index);

    void remove(int index);

    void removeHead();

    void removeTail();

    std::string printFromTail();

    std::string printFromHead();

    void removeKey(std::string key, int num);

    ListNode *find(std::string key, int num);
    
    void addBefore(std::string target, std::string key, int occurrence = 0);  
    
    void removeBefore(std::string target, int occurrence = 0);  
    
    void addAfter(std::string target, std::string key, int occurrence = 0);
    
    void removeAfter(std::string target, int occurrence = 0);
};
