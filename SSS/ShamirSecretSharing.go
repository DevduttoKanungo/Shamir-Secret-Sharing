//THIS PROGRAM IMPLEMENTS SHAMIR's SECRET SHARING
package main

//REQUIRED PACKAGES
import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
)


func main(){
	//s:="DANGER"
	s,err:=ioutil.ReadFile("src/SSS/textfile.txt") //READING A TEXT FILE
	//IF PROBLEM OCCURS DURING FILE READING THEN THE FILE READING PORTION CAN BE COMMENTED OUT AND THE STRING s DELCARED INSIDE THE PROGRAM IN LINE 17 CAN BE TAKEN OUT FROM THE COMMENT AND USED
	if err!=nil{
		panic(err)
	}
	length:=len(s) //LENGTH OF INPUT
		
	input:=257 // THE UNDERLING PRIME 
	n:=11 //ASCII OF EACH CHARACTER TO BE DIVIDED INTO 11 PARTS
	k:=5  //MINIMUM PARTS REQUIRED TO RETRIEVE EACH CHARACTER
	var arr [12]int
	var pcs [6]int
	var arr2 [6]int
	var D_pred [500]int
	
	scanner:=bufio.NewScanner(os.Stdin)
	fmt.Printf("\n\n")
	fmt.Printf("IN SHAMIR SECRET SHARING DEMO PROGRAM THE CHARACTERS OF SECRET MESSAGE CONVERETD TO ASCII WILL EACH BE DIVIDED INTO 11 PIECES.\n")
	fmt.Printf("YOU WILL BE ALLOWED TO KNOW 5 OF THEM.\n")
	fmt.Printf("YOU HAVE TO ENTER THE DESIRED INDICES FROM 1 to 11 TO KNOW THE SECRET PIECES CORRESPONDING TO THEM.\n\n")
	
	
	

	for i:=0;i<length;i++{ //THE LOOP RUNS FOR EACH CHARACTER OF THE STRING READ FROM INPUT FILE
		D:=int(s[i]) //ASCII OF EACH CHARACTER
	

		//BELOW SNIPPET IS FOR BREAKING THE ASCII OF THE CHARACTER INTO 11 PIECES BY THE ALGORITHM OF SHAMIR SECRET SHARING
		arr[0]=D
		for i:=1;i<=n;i++{
			arr[i]=poly(i,k,D,input)

		}
	
		//BELOW SNIPPET IS FOR OBTAINING 5 OF THE 11 PIECES TO REVEAL THE CHARACTERS(WE COLUD HAVE ALSO CHOSEN INDICES RANDOMLY,INSTEAD WE MANUALLY INPUT FIVE DISTINCT INDICES FOR EACH CHARACTER)
		arr2[0]=D	//arr2 ARRAY STORES THE 5 SECRET PIECES IN INDICES 1 TO 5, AND BY DEFAULT THE ZERO-TH POSITION IS SET TO ASCII OF THE CURRENT CHARACTER  
		pcs[0]=0    //pcs ARRAY STORES THE INDICES OF THE SECRETS, AND BY DEFAULT THE ZERO-TH POSITION OF THIS ARRAY IS SET TO ZERO
		fmt.Printf("Enter the indices of secret pieces:\n")
		for i:=1;i<=k;i++{
			fmt.Printf("\nEnter index %d: ",i)
			scanner.Scan()
			j,_:=strconv.ParseInt(scanner.Text(),10,64)
			fmt.Printf("\nThe secret piece chosen is: %d\n", arr[j])
			arr2[i]=arr[j]
			pcs[i]=int(j)
			fmt.Printf("The secret index chosen is: %d\n", pcs[i])
		}
	

		D_pred[i]=interpolate(0,k,arr2,pcs,input) //RETRIEVING THE SECRET CHARACTER ASCII BY LAGRANGE INTERPOLATION
	}
		fmt.Printf("\nThe secret message is: ",)
		for i:=0;i<length;i++{
			fmt.Printf("%v",string(D_pred[i])) //CONVERTING ASCII TO STRING AND PRINTING THE MESSAGE
		}
		fmt.Printf("\n\n")
	}

//THE FOLLOWING FUNCTION CALCULATES THE POLYNOMIAL MODULO GIVEN PRIME TO OBTAIN THE SECRET PIECES. INPUTS ARE INDEX x,DEGREE,SECRET,UNDERLYING PRIME
func poly(x int,n int,D int, p int) int{
	var coef [6]int
	coef[0]=D //THE CONSTANT TERM OF THE POLYNOMIAL IS THE ASCII OF THE CORRESPONDING CHARACTER

	//BELOW SNIPPET IS FOR RANDOMLY GENERATING THE COEFFICIENTS OF THE POLYNOMIAL
	for i:=1;i<=n;i++{
		RandomNum,_:=rand.Int(rand.Reader,big.NewInt(257))
		coef[i]=int(RandomNum.Int64())
		
	}

	sum:=D
	for i:=1;i<=n;i++{
		sum=(sum+coef[i]*x)%p
		x*=x
		x=x%p

	}
	return sum

}

//THE FOLLOWING FUNCTION REVEALS EACH CHARACTER ASCII BY LAGRANGE INTERPOLATION. INPUTS ARE x WHICH IS SET TO 0 DURING CALL,DEGREE,SECRET PIECES,INDICES,PRIME
func interpolate(x int,n int,ar [6]int,pcs [6]int,p int) int{
	sum:=0
	a:=1
	b:=1
	
	
	for i:=0;i<=n;i++{
		for j:=0;j<=n;j++{
			if j!=i{
				a=(a*(x-pcs[j]))%p
				b=(b*(pcs[i]-pcs[j]))%p
				if a<0{
					a=p+a
				}
				if b<0{
					b=p+b
				}

			}

		}
		c:=inverse(b,p)
		sum=(sum+a*c*ar[i])%p
		a=1
		b=1
	}
	return sum

}

//THE FOLLOWING FUNCTION FINDS THE INVERSE OF AN ELEMENT IN THE FIELD Zp WHERE p=257. COULD ALSO BE DONE BY EXTENDED EUCLIDEAN ALGORITHM
func inverse(x int, p int) int{
	var i int
	
	for i=1;i<p;i++{
		if (i*x)%p==1{
			break
		}
	}

return i


}

