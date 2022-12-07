//if
var flag = true;
if(flag) {
    console.log("判断结果是" + flag);
}

//switch
var flag1 = 2;
switch (flag1) {
    case 1:
        console.log("条件1");
        break;
    case 2:
        console.log("条件2");
        break;
    default:
        console.log("这是默认条件");
}

//for
for(var i = 0; i < 10; i++) {
    console.log("i = " + i);
}

//while
var flag2 = 0;
while (flag2 < 10) {
    console.log("flag2 = " + flag2);
    flag2++;
}

//do...while
var flag3 = 0;
do {
    console.log("flag3 = " + flag3);
    flag3++;
} while(flag3 < 10);