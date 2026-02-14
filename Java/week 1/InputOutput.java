import java.util.Scanner;
public class InputOutput {
    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        System.out.print("Enter a number:");
        int num = sc.nextInt();
        System.out.println("You enetered: " +num);
        sc.close();
    }
}