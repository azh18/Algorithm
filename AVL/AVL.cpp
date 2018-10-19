#include <cstdio>
#include<stack>
#include<vector>

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
        // delete BST node and rebalance the nodes on the trace until the subroot
        Node<T> *subroot = node, *p = this->root, *subrootFather = NULL;
        if(node == this->root){
            this->root = NULL;
            return;
        }
        stack<Node<T>* > s;
        while(p!=NULL && p!=node){
            s.push(p);
            if(node->val > p->val){
                if(p->right == node){
                    subrootFather = p;
                }
                p = p->right;
            } else {
                if(p->left == node){
                    subrootFather = p;
                }
                p = p->left;
            }
        }
        // find the leftest on the right tree
        Node<T> *p1 = subroot->right, *p1Father = subroot;
        s.push(subroot);
        if(p1 != NULL){
            if(p1->left == NULL){
                subroot->val = p1->val;
                p1Father->right = p1->right;
                p1Father->rightHeight = getHeight(p1Father->right);
                free(p1);
            } else {
                while(p1->left != NULL){
                    s.push(p1);
                    p1Father = p1;
                    p1 = p1->left;
                }
                subroot->val = p1->val;
                p1Father->left = p1->right;
                p1Father->leftHeight = getHeight(p1Father->left);
                free(p1);
            }
        } else {
            if(subroot->left != NULL){
                Node<T> *pLeft = subroot->left;
                subroot->val = pLeft->val;
                subroot->left = pLeft->left;
                subroot->right = pLeft->right;
                subroot->leftHeight = getHeight(subroot->left);
                subroot->rightHeight = getHeight(subroot->right);
                free(pLeft);
            } else {
                if(subrootFather->left == subroot){
                    subrootFather->left = NULL;
                    subrootFather->leftHeight = 0;
                } else {
                    subrootFather->right = NULL;
                    subrootFather->rightHeight = 0;
                }
                free(subroot);
                s.pop();
            }
        }
        // fix height number and fix average
        Node<T> *p2 = NULL, *p2Father = NULL, *p2New = NULL;
        while(!s.empty()){
            p2Father = s.top();
            s.pop();
            if(p2New != NULL){
                if(p2Father->left == p2){
                    p2Father->left = p2New;
                } else {
                    p2Father->right = p2New;
                }
            }
            p2 = p2Father;
            p2->leftHeight = getHeight(p2->left);
            p2->rightHeight = getHeight(p2->right);
            p2New = maintainAvg(p2);
        }
        this->root = p2New;
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
    vector<T> traverse(){
        vector<T> res;
        stack<Node<T> *> s;
        Node<T> *p = this->root;
        if(p==NULL){
            return res;
        }
        while(p->left != NULL){
            s.push(p);
            p = p->left;
        }
        res.push_back(p->val);
        while(!s.empty()){
            Node<T> *newroot = s.top();
            s.pop();
            res.push_back(newroot->val);
            p = newroot->right;
            while(p!=NULL){
                s.push(p);
                p = p->left;
            }
        }
        return res;
    } 
    
};


int main(){
    printf("hello world.\n");
    int a[10] = {1,4,2,3,5,6,8,10,7,9};
    AVLTree<int> *tree = new AVLTree<int>(a, 10);
    Node<int> *res = tree->find(3);
    tree->remove(res);
    vector<int> res1 = tree->traverse();
    printf("sorted:\n");
    for(auto it = res1.begin(); it != res1.end(); it ++){
        printf("%d ", *it);
    }
    // printf("%d, leftheight:%d, rightheight:%d\n", res->val, res->leftHeight, res->rightHeight);
    return 0;
}