#include <cstdio>
#include<stack>

using namespace std;

template<class T>
struct Node{
    T val;
    int leftHeight, rightHeight;
    Node<T> *left, *right;
};

template<class T> 
class AVLTree{


    public:
    Node<T>* root;
    
    AVLTree(){
        this->root = NULL;
    }

    AVLTree(T *input, size_t length){
        AVLTree();
        for(int i=0;i<length;i++){
            this->insert(input[i]);
        }
    }

    Node<T>* find(T val){
        Node<T> *p = this->root;
        if(p!=NULL){
            printf("%d\t", p->val);
        }
        while(p!=NULL){
            if(val == p->val){
                return p;
            } else if (val > p->val){
                p = p->right;
            } else {
                p = p->left;
            }
        }
        return NULL;
    }

    size_t getHeight(Node<T>* node){
        if(node == NULL){
            return 0;
        } else {
            return max(node->leftHeight, node->rightHeight) + 1;
        }
    }

    void insert(T val){
        Node<T> *p = this->root;
        if(this->root==NULL){
            this->root = new Node<T>{val, 0, 0, NULL, NULL};
            return;
        }

        stack<Node<T> *> trace;
        while(p->left != NULL || p->right != NULL){
            if(val > p->val){
                // insert into right subtree
                if(p->right != NULL){
                    trace.push(p);
                    p = p->right;
                } else {
                    break;
                }
            } else {
                // insert into left subtree
                if(p->left != NULL){
                    trace.push(p);
                    p = p->left;
                } else {
                    break;
                }
            }
        }

        Node<T>* newNode = new Node<T>{val,0,0, NULL, NULL};
        if(val > p->val){
            p->right = newNode;
            p->rightHeight = 1;
        } else {
            p->left = newNode;
            p->leftHeight = 1;
        }

        while(!trace.empty()){
            Node<T> *p1 = trace.top();
            trace.pop();
            if(p1->left == p){
                p1->left = maintainAvg(p);
                p1->leftHeight = getHeight(p1->left);
                p = p1;
            } else {
                p1->right = maintainAvg(p);
                p1->rightHeight = getHeight(p1->right);
                p = p1;
            }

        }

        this->root = this->maintainAvg(this->root);
        return;
    }

    void remove(Node<T> *node){
        return;
    }

    // return new root
    Node<T>* maintainAvg(Node<T> *root){
        // Node<T> *thisroot = root;
        // Node<T> *root = thisroot;
        if(root == NULL){
            return NULL;
        }
        if(root->leftHeight > root->rightHeight + 1){
            Node<T> *pLeft = root->left;
            if(pLeft->leftHeight > pLeft->rightHeight){
                // LL
                Node<T> *LR = pLeft->right;
                // this->root = pLeft;
                pLeft->right = root;
                root->left = LR;
                root->leftHeight = getHeight(root->left);
                pLeft->rightHeight = getHeight(pLeft->right);
                // new root is pLeft
                return pLeft;
            } else {
                // LR
                Node<T> *RL = pLeft->right->left;
                Node<T> *newroot = pLeft->right;
                root->left = newroot;
                newroot->left = pLeft;
                pLeft->right = RL;
                pLeft->rightHeight = getHeight(pLeft->right);
                newroot->leftHeight = getHeight(newroot->left);
                root->leftHeight = getHeight(root->left);
                newroot = maintainAvg(root); // To LL
                return newroot;
            }
        } else if (root->rightHeight > root->leftHeight + 1){
            Node<T> *pRight = root->right;
            if(pRight->leftHeight > pRight->rightHeight){
                // RL
                Node<T> *RL = pRight->left;
                Node<T> *RLR = RL->right;
                root->right = RL;
                RL->right = pRight;
                pRight->left = RLR;
                pRight->leftHeight = getHeight(pRight->left);
                RL->rightHeight = getHeight(RL->right);
                root->rightHeight = getHeight(root->right);
                Node<T> *newroot = maintainAvg(root); // to RR
                return newroot;
            } else {
                // RR
                Node<T> *RL = pRight->left;
                // this->root = pRight;
                pRight->left = root;
                root->right = RL;
                root->rightHeight = getHeight(root->right);
                pRight->leftHeight = getHeight(pRight->left);
                // new root is pRight
                return pRight;
            }
        }
        return root;
    }
};


int main(){
    printf("hello world.\n");
    int a[10] = {1,4,2,3,5,6,8,10,7,9};
    AVLTree<int> *tree = new AVLTree<int>(a, 10);
    Node<int> *res = tree->find(5);
    printf("%d, leftheight:%d, rightheight:%d\n", res->val, res->leftHeight, res->rightHeight);
    return 0;
}