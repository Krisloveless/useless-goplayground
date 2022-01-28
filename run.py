f = open("Book1.csv", "r")
x = f.readline().rstrip().split(',')
y = f.readline().rstrip().split(',')
f.close()
interval = int(len(x) / 9)
print(interval)
t = 0
string1 = "var x = []int{"
string2 = "var y = []int{"
arr1 = []
arr2 = []
while t < len(x):
	arr1.append(str(int(float(x[t])))) 
	arr2.append(str(int(float(y[t]))))
	t += interval

arr1 = arr1[::-1] + arr1
arr2 = arr2[::-1] + arr2
arr1 += arr1 + arr1
arr2 += arr2 + arr2
print(len(arr1))
string1 += ",".join(arr1) + "}" 
string2 += ",".join(arr2) + "}"
print(string1)
print(string2)
