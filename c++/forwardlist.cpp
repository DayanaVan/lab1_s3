#include "forwardlist.h"

ForwardListNode::ForwardListNode(std::string key)
{
    this->key = key;
    this->next = nullptr;
}

ForwardList::ForwardList()
{
    this->head = nullptr;
}

// Вставка по индексу
void ForwardList::insert(std::string key, int index)
{
    if(index == 0)
    {
        addHead(key);
        return;
    }
    
    ForwardListNode *node = this->head;
    for(int i = 0; i < index - 1; i++)
    {
        if(node == nullptr)
            break;
        node = node->next;
    }
    if(node == nullptr)
        return;
    ForwardListNode *newNode = new ForwardListNode(key);
    newNode->next = node->next;
    node->next = newNode;

}

// Удаление по индексу
void ForwardList::remove(int index)
{
    if(index == 0)
    {
        removeHead();
        return;
    }

    ForwardListNode *node = this->head;
    for(int i = 0; i < index - 1; i++)
    {
        if(node == nullptr)
            break;
        node = node->next;
    }
    if(node == nullptr)
        return;
    ForwardListNode *toDelete = node->next;
    if(toDelete != nullptr)
    {
        node->next = toDelete->next;
        delete toDelete;
    }
}

// Добавление в хвост
void ForwardList::addTail(std::string key)
{
    if(this->head == nullptr)
    {
        this->head = new ForwardListNode(key);
        return;
    }
    ForwardListNode *node = this->head;
    while(node->next != nullptr)
        node = node->next;
    node->next = new ForwardListNode(key);
}

// Добавление в голову
void ForwardList::addHead(std::string key)
{
    if(this->head == nullptr)
    {
        this->head = new ForwardListNode(key);
        return;
    }
    ForwardListNode *node = new ForwardListNode(key);
    node->next = this->head;
    this->head = node;
}

// Удаление хвоста
void ForwardList::removeTail()
{
    if(this->head == nullptr)
        return;

    ForwardListNode *node = this->head;
    while(node->next->next != nullptr)
        node = node->next;
    delete node->next;
    node->next = nullptr;
}

// Удаление головы
void ForwardList::removeHead()
{
    if(this->head == nullptr)
        return;

    ForwardListNode *toDelete = this->head;
    this->head = toDelete->next;
    delete toDelete;
}

// Вывод с головы
std::string ForwardList::printFromHead()
{
    ForwardListNode *node = this->head;
    if(node == nullptr)
        return "";
    std::string a = "";
    while(node->next != nullptr)
    {
        a += node->key + ' ';
        node = node->next;
    }
    a += node->key;
    return a;
}

// Вывод с хвоста
void ForwardList::printFromTailRec(std::string &s, ForwardListNode *node)
{
    if(node == nullptr)
        return;
    printFromTailRec(s, node->next);
    s += node->key + " ";
}

std::string ForwardList::printFromTail()
{
    std::string s = "";
    printFromTailRec(s, this->head);
    return s;
}

// Удаление элемента по индексу
void ForwardList::removeKey(std::string key, int num)
{
    if(num < 0 || this->head == nullptr)
        return;
    int n = 0;
    ForwardListNode *node = this->head;
    if(node->key == key && num == 0)
    {
        removeHead();
        return;
    }
    while(node->next != nullptr)
    {
        if(node->next->key == key)
        {
            if(n == num)
            {
                ForwardListNode *toDelete = node->next;
                node->next = toDelete->next;
                delete toDelete;
                return;
            }
            n++;
        }
        node = node->next;
    }
}

// Нахождение элемента
ForwardListNode *ForwardList::find(std::string key, int num)
{
    if(num < 0)
        return nullptr;
    int n = 0;
    ForwardListNode *node = this->head;
    while(node != nullptr)
    {
        if(node->key == key)
        {
            if(n == num)
                return node;
            n++;
        }
        node = node->next;
    }
    return nullptr;
}

// Добавление до элемента 
void ForwardList::addBefore(std::string target, std::string key, int occurrence)
{
    if (occurrence < 0 || this->head == nullptr)
        return;

    int n = 0;
    ForwardListNode* prev = nullptr;
    ForwardListNode* current = this->head;

    if (current->key == target && occurrence == 0)
    {
        addHead(key);
        return;
    }

    while (current != nullptr)
    {
        if (current->key == target)
        {
            if (n == occurrence)
            {
                if (prev == nullptr) // голова
                {
                    addHead(key);
                }
                else
                {
                    ForwardListNode* newNode = new ForwardListNode(key);
                    newNode->next = current;
                    prev->next = newNode;
                }
                return;
            }
            n++;
        }
        prev = current;
        current = current->next;
    }
}

// Удаление до элемента
void ForwardList::removeBefore(std::string target, int occurrence)
{
    if (occurrence < 0 || this->head == nullptr || this->head->next == nullptr)
        return;

    int n = 0;
    ForwardListNode* prevPrev = nullptr;
    ForwardListNode* prev = this->head;
    ForwardListNode* current = this->head->next;

    if (prev->key == target && occurrence == 0)
    {
        return;
    }

    while (current != nullptr)
    {
        if (current->key == target)
        {
            if (n == occurrence)
            {
                if (prevPrev == nullptr) // голова
                {
                    removeHead();
                }
                else
                {
                    prevPrev->next = current;
                    delete prev;
                }
                return;
            }
            n++;
        }
        prevPrev = prev;
        prev = current;
        current = current->next;
    }
}

// Добавление после элемента 
void ForwardList::addAfter(std::string target, std::string key, int occurrence)
{
    if (occurrence < 0)
        return;

    int n = 0;
    ForwardListNode* node = this->head;

    while (node != nullptr)
    {
        if (node->key == target)
        {
            if (n == occurrence)
            {
                ForwardListNode* newNode = new ForwardListNode(key);
                newNode->next = node->next;
                node->next = newNode;
                return;
            }
            n++;
        }
        node = node->next;
    }
}

// Удаление после элемента
void ForwardList::removeAfter(std::string target, int occurrence)
{
    if (occurrence < 0 || this->head == nullptr)
        return;

    int n = 0;
    ForwardListNode* node = this->head;

    while (node != nullptr)
    {
        if (node->key == target)
        {
            if (n == occurrence)
            {
                if (node->next != nullptr)
                {
                    ForwardListNode* toDelete = node->next;
                    node->next = toDelete->next;
                    delete toDelete;
                }
                return;
            }
            n++;
        }
        node = node->next;
    }
}
