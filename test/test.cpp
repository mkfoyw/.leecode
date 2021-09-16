int a[20];  //保存拆分的数位

ll dp[20][state];   //不同题目状态不同

ll dfs(int pos //枚举位置, /*state变量*/  ,bool lead/*前导零*/,  bool limit /*数位上界变量*/)  //不是每个题都要判断前导零
{

//递归边界，既然是按位枚举，最低位是0，那么pos==-1说明这个数我枚举完了
if(pos==-1) return 1;/*这里一般返回1，表示你枚举的这个数是合法的，那么这里就需要你在枚举时必须每一位都要满足题目条件，也就是说当前枚举到pos位，一定要保证前面已经枚举的数位是合法的。不过具体题目不同或者写法不同的话不一定要返回1 */

//第二个就是记忆化(在此前可能不同题目还能有一些剪枝)
if(!limit && !lead && dp[pos][state]!=-1) return dp[pos][state];

//判断是否存在枚举限制
int up=limit?a[pos]:9;

ll ans=0;
//开始计数
for(int i=0;i<=up;i++)
{if() ...
else if()...
ans+=dfs(pos-1,/*状态转移*/,lead && i==0,limit && i==a[pos]) //最后两个变量传参都是这样写的
}

//计算完，记录状态, 不要记录有限制的状态
if(!limit && !lead) dp[pos][state]=ans;
return ans;
}
