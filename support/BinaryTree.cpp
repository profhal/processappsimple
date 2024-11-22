#include <iostream>
#include <string>
#include "BinaryTree.h"

using namespace std;

BinaryTree::BinaryTree() {

    root = nullptr;

    nodeCount = 0;

}



BinaryTree::BinaryTree(const BinaryTree &tree) {

    string *nodeValues = new string[tree.getNodeCount()]; 

    storeInOrder(tree.root, nodeValues, nodeCount);

    root = buildBalancedTree(nodeValues, 0, nodeCount-1);

}



BinaryTree::~BinaryTree() {

    destroySubtree(root);

}



void BinaryTree::insert (TreeNode *&node, TreeNode *&newNode) {

    if (node == nullptr) {

        node = newNode;

    } else if (newNode->value <= node->value) {
    
        insert(node->left, newNode);
    
    } else {
    
        insert(node->right, newNode);
        
    }

}



void BinaryTree::destroySubtree (TreeNode *node) {
    
    if (node != nullptr) {
    
    	if (node->left != nullptr) {
    		destroySubtree(node->left);    		
    	}
        
        if (node->right != nullptr) {
    		destroySubtree(node->right);
    	}
    	
    	delete node;
    
    }
    
}



void BinaryTree::deleteNode (string value, TreeNode *&node) {
    if (value < node->value) {
    	deleteNode(value, node->left);
    } else if (value > node->value) {
    	deleteNode(value, node->right);
    } else {
    	makeDeletion(node);
    }
}



void BinaryTree::makeDeletion (TreeNode *&node) {
    
    TreeNode *tempNode = nullptr;
    
    // Quick exit
    if (node == nullptr) {
 		return;    
    } else if (node->right == nullptr) { // no right child, reattach left
    	tempNode = node;
    	node = node->left;
    	delete tempNode;
    } else if (node->left == nullptr) { // no left child, reattach right
    	tempNode = node;
    	node = node->right;
    	delete tempNode;
    } else {					  // there are two children
    	
    	// Move a node to the right
    	//
    	tempNode = node->right;
    	
    	// Go the far left side of tempNode
    	//
    	while (tempNode->left != nullptr) {
    		tempNode = tempNode->left;
    	}
    	
    	// Reattach the left subtree
    	//
    	tempNode->left = node->left;
    	tempNode = node;
    	
    	// Reattach right subtree
    	//
    	node = node->right;
    	
    	delete tempNode;
    	
    }


}



void BinaryTree::displayInorder (TreeNode *node)  const {
    if (node != nullptr) {
        displayInorder(node->left);
        cout << node->value << endl;
        displayInorder(node->right);
    }    
}



void BinaryTree::displayPreorder (TreeNode *node) const {
    if (node != nullptr) {
        cout << node->value << endl;
        displayPreorder(node->left);
        displayPreorder(node->right);
    }    
}



void BinaryTree::displayPostorder(TreeNode *node) const {
    if (node != nullptr) {
        displayPostorder(node->left);
        displayPostorder(node->right);
        cout << node->value << endl;
    }    
}



void BinaryTree::storeInOrder(TreeNode* root, string nodeValues[], int &nodeCount) {

    if (root != nullptr) {

        storeInOrder(root->left, nodeValues, nodeCount);

        nodeValues[nodeCount++] = root->value;

        storeInOrder(root->right, nodeValues, nodeCount);

    } 

}



BinaryTree::TreeNode* BinaryTree::buildBalancedTree(string nodeValues[], int start, int end) {

    TreeNode* result = nullptr;
    int mid;    

    if (start <= end) {

        mid = (start + end) / 2;

        result = new TreeNode(nodeValues[mid], 
                              buildBalancedTree(nodeValues, start, mid-1), 
                              buildBalancedTree(nodeValues, mid + 1, end) );

    }

    return result;

}



void BinaryTree::balanceTree() {

    string *nodeValues = new string[nodeCount];
    int nodeCount = 0;

    storeInOrder(root, nodeValues, nodeCount);

    delete root;

    root = buildBalancedTree(nodeValues, 0, nodeCount -1);

}



int BinaryTree::getNodeCount() const {

    return nodeCount;

}



int BinaryTree::maxDepth(TreeNode* root) const {

    int leftDepth;
    int rightDepth;

    if (root != nullptr) {

        leftDepth = maxDepth(root->left);
        rightDepth = maxDepth(root->right);

        return max(leftDepth, rightDepth) + 1;

    } else {

        return 0;

    }
    
}



int BinaryTree::getHeight() const {

    return maxDepth(root);

}



void BinaryTree::insert(string value) {

    TreeNode *newNode;
    
    newNode = new TreeNode(value, nullptr, nullptr);
    
    insert(root, newNode);

    ++nodeCount;

}



bool BinaryTree::has(string value) const {
    TreeNode *node = root;
    
    bool found = false;
    
    while (!found && node != nullptr) {
    
    	if (node->value == value) {
    		found = true;
    	} else if (value < node->value) {
    		node = node->left;
    	} else {
    		node = node->right;
    	}
    
    }

	return found;

}



void BinaryTree::remove(string value) {
    deleteNode(value, root);
    --nodeCount;
}



void BinaryTree::displayInorder() const {
    displayInorder(root);
}



void BinaryTree::displayPreorder() const {
    displayPreorder(root);
}



void BinaryTree::displayPostorder() const {
    displayPostorder(root);
}



string* BinaryTree::getAscendingArray() {

    string* nodeValuesInOrder = new string[nodeCount];
    int nodeValueCount = 0;

    storeInOrder(root, nodeValuesInOrder, nodeValueCount);

    return nodeValuesInOrder;
    
}



BinaryTree& BinaryTree::operator=(const BinaryTree &tree) {

    string *nodeValues = new string[tree.getNodeCount()]; 

    storeInOrder(tree.root, nodeValues, nodeCount);

    destroySubtree(root);

    root = buildBalancedTree(nodeValues, 0, nodeCount-1);  

    return *this;

}