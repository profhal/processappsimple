#ifndef BINARY_TREE_H
#define BINARY_TREE_H

#include<string>

using namespace std;


class BinaryTree {

private:

    class TreeNode {
    public:
        string   value;
        TreeNode *left;
        TreeNode *right;
    
        TreeNode(string initValue, TreeNode* initLeft, TreeNode* initRight) {
            value = initValue;
            left = initLeft;
            right = initRight;
        }
    };
    
    TreeNode *root;

    int nodeCount;

    
    void insert         (TreeNode *&node, TreeNode *&newNode);
    void destroySubtree (TreeNode *node);
    void deleteNode     (string value, TreeNode *&node);
    void makeDeletion   (TreeNode *&node);
    
    void displayInorder  (TreeNode *node) const;
    void displayPreorder (TreeNode *node) const;
    void displayPostorder(TreeNode *node) const;

    void storeInOrder(TreeNode* root, string nodeValues[], int &nodeCount);
    TreeNode* buildBalancedTree(string nodeValues[], int start, int end);

    int maxDepth(TreeNode* root) const;
    
public:

    BinaryTree();

    BinaryTree(const BinaryTree &tree);
    
    virtual ~BinaryTree();

    virtual void insert(string value);
    virtual bool has(string value) const;
    virtual void remove(string value);
    
    virtual void displayInorder() const;
    virtual void displayPreorder() const;
    virtual void displayPostorder() const;

    virtual void balanceTree();
    virtual int getNodeCount() const;
    virtual int getHeight() const;

    virtual string* getAscendingArray();
    
    virtual BinaryTree& operator=(const BinaryTree &rhs);

};


#endif