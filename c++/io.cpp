#include "io.h"

int readArray(std::ifstream &ifs, Array *arr)
{
    std::string a;
    while(ifs >> a)
    {
        arr->add(a);
    }
    return 0;
}

void writeArray(std::ofstream &ofs, Array *arr)
{
    ofs << arr->print();
}

int readForwardList(std::ifstream &ifs, ForwardList *fl)
{
    std::string a;
    ForwardListNode *tail;
    while(ifs >> a)
    {
        if(fl->head == nullptr)
        {
            fl->addHead(a);
            tail = fl->head;
        }
        else
        {
            tail->next = new ForwardListNode(a);
            tail = tail->next;
        }
    }
    return 0;
}

void writeForwardList(std::ofstream &ofs, ForwardList *fl)
{
    ofs << fl->printFromHead();
}

int readList(std::ifstream &ifs, List *l)
{
    std::string a;
    while(ifs >> a)
    {
        l->addTail(a);
    }
    return 0;
}

void writeList(std::ofstream &ofs, List *l)
{
    ofs << l->printFromHead();
}

int readStack(std::ifstream &ifs, Stack *s)
{
    std::string a;
    while(ifs >> a)
    {
        s->push(a);
    }
    return 0;
}

void writeStackRec(std::ofstream &ofs, StackNode *sn)
{
    if(sn == nullptr)
        return;
    writeStackRec(ofs, sn->prev);
    ofs << sn->key << " ";
}

void writeStack(std::ofstream &ofs, Stack *s)
{
    writeStackRec(ofs, s->top);
}

int readQueue(std::ifstream &ifs, Queue *q)
{
    std::string a;
    while(ifs >> a)
    {
        q->push(a);
    }
    return 0;
}

void writeQueue(std::ofstream &ofs, Queue *q)
{
    ofs << q->print();
}

// Чтение узла дерева из файла
int readTreeNode(std::ifstream &ifs, TreeNode *&node)
{
    std::string token;
    if (!(ifs >> token)) {
        return 1; // Конец файла
    }
    
    if (token == "null" || token == "00") {
        node = nullptr;
    } else if (isNumber(token)) {
        node = new TreeNode(std::stoi(token));
    } else {
        return 2; // Некорректный формат
    }
    
    return 0;
}

// Чтение дерева из файла (BFS формат)
int readTree(std::ifstream &ifs, Tree *t)
{
    // Проверяем, пустой ли файл
    if (ifs.peek() == std::ifstream::traits_type::eof()) {
        t->root = nullptr;
        return 0;
    }
    
    std::queue<TreeNode**> nodesQueue;
    nodesQueue.push(&(t->root));
    
    std::string token;
    while (!nodesQueue.empty() && ifs >> token) {
        TreeNode** current = nodesQueue.front();
        nodesQueue.pop();
        
        if (token == "null" || token == "00") {
            *current = nullptr;
        } else if (isNumber(token)) {
            *current = new TreeNode(std::stoi(token));
            // Добавляем указатели на детей в очередь
            nodesQueue.push(&((*current)->left));
            nodesQueue.push(&((*current)->right));
        } else {
            return 1; // Некорректный формат
        }
    }
    
    return 0;
}

// Запись дерева в файл (BFS формат)
void writeTree(std::ofstream &ofs, Tree *t)
{
    if (t->root == nullptr) {
        ofs << "";
        return;
    }
    
    std::queue<TreeNode*> q;
    q.push(t->root);
    
    bool first = true;
    while (!q.empty()) {
        TreeNode* node = q.front();
        q.pop();
        
        if (!first) {
            ofs << " ";
        }
        first = false;
        
        if (node == nullptr) {
            ofs << "null";
        } else {
            ofs << node->key;
            q.push(node->left);
            q.push(node->right);
        }
    }
}

bool isNumber(std::string s)
{
    if(!(s[0] == '-' || (s[0] >= '0' && s[0] <= '9')))
        return false;

    for(int i = 1; i < s.size(); i++)
    {
        if(s[i] < '0' || s[i] > '9')
            return false;
    }
    return true;
}
