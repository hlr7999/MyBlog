import axios from 'axios'

var baseUrl = "http://localhost:2333/"

// login
const UserLogin = function(data) {
    return axios.post(baseUrl + "login", {
        username: data.username,
        password: data.pass
    })
}

// register
const UserRegister = function(data) {
    return axios.post(baseUrl + "register", {
        username: data.username,
        password: data.pass,
        email: data.email
    })
}

// aboutMe
const AboutMe = function() {
    return axios.get(baseUrl + "aboutMe")
}


// //用户退出
// const LoginOut = (token,callback) =>{
//     let url = portUrl + 'login/LoginOut?token='+token;
//     axios.get(url).then(num => {
//             callback && callback(num.data);
//     })
// }

// //文章分类查询
// const ArtClassData = (callback) => {
//     if(sessionStorage.getItem('classList')){
//         var data = JSON.parse(sessionStorage.getItem('classList'));
//         callback && callback(data)
//     }else{
//         let url = portUrl + 'article/ArtClassData';
//         axios.get(url).then(num => {
//             // console.log(num);
//             if(num.data.code==1001){
//                 sessionStorage.setItem('classList',JSON.stringify(num.data.data));
//                 callback && callback(num.data.data)
//             }else{
//                 alert("查询失败")
//             }
//         })
//     }
// }

// //实验室 列表项目
// const navMenList  = (callback) => {
//     if(sessionStorage.getItem('navMenList')){
//         var data = JSON.parse(sessionStorage.getItem('navMenList'));
//         callback && callback(data)
//     }else{
//         let url = portUrl + 'nav/navMenList';
//         axios.get(url).then(num => {
//             // console.log(num);
//             if(num.data.code==1001){
//                 sessionStorage.setItem('navMenList',JSON.stringify(num.data.data));
//                 callback && callback(num.data.data)
//             }else{
//                 alert("查询失败")
//             }
//         })
//     }
// }

// //查询文章列表
// const ShowArticleAll = (artId,cateId,articleName,level,callback) =>{
//     if(level == 1){
//         var url = portUrl + 'nav/ActiveClassAllData?art_id='+artId+'&cate_id='+cateId+'&article_name='+articleName;
//     }else{
//         var url = portUrl + 'article/ShowArticleAll?art_id='+artId+'&cate_id='+cateId+'&article_name='+articleName;
//     }
//     axios.get(url).then(num => {
//             callback && callback(num.data);
//     })
// }

// //查询文章详情
// const getArticleInfo = (artId,userId,callback) =>{
//     let url = portUrl + 'article/getArticleInfo?art_id='+artId+'&user_id='+userId;
//     axios.get(url).then(num => {
//         if(num.data.code==1001){
//             callback && callback(num.data.data);
//         }else{
//             alert("查询失败");
//         }
//     })
// }

// //查询浏览量最多的10篇文章数据
// const ShowBrowseCount = (callback) =>{
//     let url = portUrl + 'article/ShowBrowseCount';
//     axios.get(url).then(num => {
//         if(num.data.code==1001){
//             callback && callback(num.data.data);
//         }else if(num.data.code==1005){
//             return;
//         }else{
//             alert("查询失败");
//         }
//     })
// }

// //查询文章评论量最大的10篇文章
// const ShowArtCommentCount = (callback) =>{
//     let url = portUrl + 'article/ShowArtCommentCount';
//     axios.get(url).then(num => {
//         if(num.data.code==1001){
//             callback && callback(num.data.data);
//         }else if(num.data.code==1005){
//             return;
//         }else{
//             alert("查询失败");
//         }
//     })
// }
// //查询文章评论数据
// const ArticleComment = (artId,commentId,callback) =>{
//     let url = portUrl + 'comment/ArticleComment?art_id='+artId+'&comment_id='+commentId;
//     axios.get(url).then(num => {
//             callback && callback(num.data);
//     })
// }

// //查询其他评论数据
// const OtherComment = (leaveId,commentId,callback) =>{//分类类型ID（1：赞赏 2：友情链接 3：留言板 4：关于我）
//     let url = portUrl + 'comment/OtherComment?leave_id='+leaveId+'&comment_id='+commentId;
//     axios.get(url).then(num => {
//         callback && callback(num.data);
//     })
// }

// //文章评论
// const setArticleComment = (content,user_id,article_id,leave_pid,pid,callback) =>{
//     let url = portUrl + 'comment/setArticleComment?content='+content+'&user_id='+user_id+'&article_id='+article_id+'&leave_pid='+leave_pid+'&pid='+pid;
//     axios.get(url).then(num => {
//             callback && callback(num.data);
//     })
// }

// //其他评论
// const setOuthComment = (content,user_id,article_id,leave_id,leave_pid,pid,callback) =>{
//     let url = portUrl + 'comment/setOuthComment?content='+content+'&user_id='+user_id+'&article_id='+article_id+'&leave_id='+leave_id+'&leave_pid='+leave_pid+'&pid='+pid;
//     axios.get(url).then(num => {
//             callback && callback(num.data);
//     })
// }

// //查询网址点赞总数
// const showLikeData = (callback) =>{
//     let url = portUrl + 'outh/showLikeData';
//     axios.get(url).then(num => {
//         if(num.data.code==1001){
//             // console.log(num.data,parseInt(num.data));
//             callback && callback(num.data.data);
//         }else{
//             alert("查询失败");
//         }
//     })
// }

// //点赞功能修改
// const GetLike = (like_num,callback) =>{
//     let url = portUrl + 'outh/GetLike?like_num='+like_num;
//     axios.get(url).then(num => {
//         if(num.data.code==1001){
//             callback && callback(num.data.msg);
//         }else{
//             alert("点赞失败");
//         }
//     })
// }

// //查询友情链接数据
// const FriendUrlData = (callback) =>{
//     let url = portUrl + 'outh/FriendUrlData';
//     axios.get(url).then(num => {
//         if(num.data.code==1001){
//             callback && callback(num.data.data);
//         }else if(num.data.code==1005){
//             return;
//         }else{
//             alert("查询失败");
//         }
//     })
// }

// //查询关于我
// const AboutMeData = (callback) =>{
//     if(sessionStorage.getItem('AboutMeData')){
//         var data = JSON.parse(sessionStorage.getItem('AboutMeData'));
//         callback && callback(data)
//     }else{
//         let url = portUrl + 'outh/AboutMeData';
//         axios.get(url).then(num => {
//             if(num.data.code==1001){
//                 sessionStorage.setItem('AboutMeData',JSON.stringify(num.data.data));
//                 callback && callback(num.data.data);
//             }else if(num.data.code==1005){
//                 return;
//             }else{
//                 alert("查询失败");
//             }
//         })
//     }
// }

// //文章点击收藏 点击喜欢
// const getArtLikeCollect = (userId,artId,islike,callback) =>{
//     var url = '';
//     if(islike==1){
//         url = portUrl + 'article/getArtLike?user_id='+userId+'&art_id='+artId;
//     }else{
//         url = portUrl + 'article/getArtCollect?user_id='+userId+'&art_id='+artId;
//     }
//     axios.get(url).then(num => {
//         if(num.data.code==1001){
//             callback && callback(num.data.msg);
//         }else{
//             alert("查询失败");
//         }
//     })
// }

// //查询赞赏数据
// const AdmireData = (callback) => {
//     let url = portUrl + 'outh/AdmireData';
//     axios.get(url).then(num => {
//         if(num.data.code==1001){
//             callback && callback(num.data);
//         }else{
//             alert("查询失败");
//         }
//     })
// }

// //查询用户喜欢列表,查询用户收藏列表
// const getLikeCollectList = (userId,artId,articleName,islike,callback)=>{
//     var url = '';
//     if(islike==1){
//          url = portUrl + 'article/getLikeList?user_id='+userId+'&art_id='+artId+'&article_name='+articleName;
//     }else{
//          url = portUrl + 'article/getCollectList?user_id='+userId+'&art_id='+artId+'&article_name='+articleName;
//     }
//     axios.get(url).then(num => {
//             callback && callback(num.data);
//     })
// }

// //查询用户信息
// const getUserInfo = (userId,callback)=>{
//     let url = portUrl + 'Userinfo/getUserInfo?user_id='+userId;
//     axios.get(url).then(num => {
//         if(num.data.code==1001){
//             callback && callback(num.data);
//         }else{
//             alert("查询失败");
//         }
//     })
// }
// //修改用户信息
// const UserInfoSave = (obj,callback) =>{
//     let url = portUrl + 'Userinfo/UserInfoSave';
//     var data = {
//         'username':obj.username,
//         'user_img':obj.avatar,
//         'email':obj.email,
//         'sex':obj.sex,
//         'friend_start':obj.state,
//         'user_id':obj.user_id,
//         'frie_name':obj.name,
//         'frie_url':obj.url,
//         'frie_description':obj.description,
//         'friend_img':obj.image,
//         'label':obj.label,
//         'head_start':obj.head_start,
//         'logo_start':obj.logo_start
//     };
//     // console.log(data);
//     axios.get(url,{params:data}).then(num => {
//         if(num.data.code==1001){
//             callback && callback(num.data.msg);
//         }else{
//             alert("保存失败");
//         }
//     })
// }

// //初始化时间
// const initDate = (oldDate,full) => {
//     var odate = new Date(oldDate);
//     var year =  odate.getFullYear();
//     var month = odate.getMonth()<9? '0' + (odate.getMonth()+1) : odate.getMonth()+1;
//     var date = odate.getDate()<10? '0'+odate.getDate() : odate.getDate();
//     if(full=='all'){
//         var t = oldDate.split(" ")[0];
//         // console.log(oldDate,t.split('-')[0],t.split('-')[1],t.split('-')[2]);
//         return t.split('-')[0]+'年'+t.split('-')[1]+'月'+t.split('-')[2]+'日';
//     }else if(full=='year'){
//         return year
//     }else if(full== 'month'){
//         return odate.getMonth()+1
//     }else if(full == 'date'){
//         return date
//     }else if(full== 'newDate'){
//         return year+'年'+month+'月'+date+'日';
//     }
// }

// //获取主题信息
// const changeTheme = (callback) => {
//     if(sessionStorage.getItem('changeThemeObj')){
//         var data = JSON.parse(sessionStorage.getItem('changeThemeObj'));
//         callback && callback(data)
//     }else{
//         let url = portUrl + 'outh/ThemeMy';
//         axios.get(url).then(num => {
//             if(num.data.code==1001){
//                 sessionStorage.setItem('changeThemeObj',JSON.stringify(num.data.data))
//                 callback && callback(num.data.data);
//             }else{
//                 alert("查询失败");
//             }
//         })
//     }
// }

export {
    UserLogin,
    UserRegister,
    AboutMe
}
