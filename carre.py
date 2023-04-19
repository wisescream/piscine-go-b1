import tkinter as tk
def somme_ligne(M,i):
    n=len(M)
    s=0
    for j in range(n):
        s+=M[i][j]
     
    return s

def somme_colonne(M,j):
    n=len(M)
    s=0
    for i in range(n):
        s+=M[i][j]
     
    return s

def somme_diagonale(M):
    n=len(M)
    s=0
    for i in range(n):
        s+=M[i][i]
     
    return s

def somme_antidiagonale(M):
    n=len(M)
    s=0
    for i in range(n):
        s+=M[i][n-i-1]
     
    return s

def carre_magique(C):
    n=len(C)
    ref=somme_ligne(C,0)
    for i in range(1,n):
        if ref != somme_ligne(C,i):
            return False
     
    for j in range(n):
        if ref!= somme_colonne(C,j):
            return False
     
    if somme_diagonale(C)!=ref or somme_antidiagonale(C)!=ref:
        return False
     
    return True

def magique_normal(C):
    if carre_magique(C)==False:
        return False
 
    n=len(C)
    etat=[0]* (n**2)
    for i in range(n):
        for j in range(n):
            if C[i][j]<=(n**2) and etat[C[i][j]-1]==0:
                etat[C[i][j]-1]=1
            else:
                return False
     
    return True

def matrice_nulle(n):
    return [[0]*n for i in range(n)]

def crée_carré(n):
    C=matrice_nulle(n)
    C[0][n//2]=1
    i,j=0,n//2
    it=1
    p1,p2=0,0
    while it<n**2:
        p1,p2=i,j
        i-=1
        j+=1
 
        if i<0:
            i=n-1
         
        if j>=n:
            j=0
         
        if C[i][j]!=0:
            i,j=p1+1,p2
 
        it+=1
        C[i][j]=it
     
    return C

def constante(n):
    return (n**2+1)*(n//2) +(n**2-(n+1)*(n//2))

def constante(n):
    return (n**2+1)*(n//2) +(n**2-(n+1)*(n//2))

if __name__ == "__main__":
    n = int(input("Donner la taille du carré magique:"))
    while n % 2 == 0:
        print("Veuillez entrer un nombre impair.")
        n = int(input("Donner la taille du carré magique:"))

    magic_square = crée_carré(n)
    print("Le carré magique pour  n =", n)
    for row in magic_square:
        print(*row,"\t")
    print("Somme de chaque ligne ou colonnes", constante(n))
