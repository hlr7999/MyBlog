import service from './request'

const loginServer = (data, callback) => {
    var params = new URLSearchParams();
    params.append('value', data);
    service.post(
        'server/login.php',
        params
    ).then(res => {
        callback && callback(res);
    }).catch(error => {
        console.log(error);
    });
}

export {loginServer};