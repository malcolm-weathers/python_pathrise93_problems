# Given a list of integers, find the maximum sum with no two consecutive
# integers being added.

def maxsum(x):
    if x == []:
        return []
    if len(x) == 1:
        return [x[0]]
    o1 = [x[0]] + maxsum(x[2:])
    o2 = [x[1]] + maxsum(x[3:])
    if sum(o1) > sum(o2):
        return o1
    return o2

def maxsum_dyn(x):
    dyn = []
    for y in range(len(x)):
        dyn.append((-1, []))
    dyn[-1] = (x[-1], [x[-1]])
    dyn[-2] = (x[-2], [x[-2]])
    dyn[-3] = (x[-1] + x[-3], [x[-1], x[-3]])

    for y in range(-4, -len(x)-1, -1):
        best = max(dyn[y+2][0], dyn[y+3][0])
        if best == dyn[y+2][0]:
            dyn[y] = (dyn[y+2][0]+x[y], dyn[y+2][1]+[x[y]])
        else:
            dyn[y] = (dyn[y+3][0]+x[y], dyn[y+3][1]+[x[y]])
    best = list(reversed(dyn[dyn.index(max(dyn))][1]))
    return best
    

if __name__ == '__main__':
    print(maxsum([1,2,3]))
    print(maxsum_dyn([1,2,3]))
    print(maxsum([5,1,2,6]))
    print(maxsum_dyn([5,1,2,6]))
    print(maxsum([5,1,2,6,20,2]))
    print(maxsum_dyn([5,1,2,6,20,2]))
