#define wait(s) atomic{s>0 -> s--}
#define signal(s) s++

#define N 10

int mutex = 1
int X[N] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
int Y[N] = {5, 7, 9, 11, 13, 15, 17, 19, 21, 23}
int sumX = 0, sumY = 0, sumXY = 0, sumXX = 0
int i;

active proctype LinearRegression() {

    for (i : 0 .. N - 1) {
        wait(mutex)
        sumX = sumX + X[i]
        sumY = sumY + Y[i]
        sumXY = sumXY + (X[i] * Y[i])
        sumXX = sumXX + (X[i] * X[i])
        signal(mutex)
    }

    int m = (N * sumXY - sumX * sumY) / (N * sumXX - sumX * sumX)
    int b = (sumY - m * sumX) / N

    printf("Coeficiente m: %d, Coeficiente b: %d\n", m, b)
}
