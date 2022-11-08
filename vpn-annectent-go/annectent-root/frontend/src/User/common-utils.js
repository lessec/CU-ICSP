export let addUser = async (name, email, type) => {
    let url = 'http://localhost:5555/api/v1/user/add';
    
    const params = {
        Name: name,
        Email: email,
        Type: type
    }

    console.log(JSON.stringify(params))
    const options = {
        mode: "no-cors",
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded",
        },
        body: JSON.stringify(params)
    }

    return await fetch(url, options).then(res => {
        console.log(res);
    }).catch(err => {
        console.log(err);
    });
}

export let connectGroup = async (uid, gid) => {
    let url = `http://localhost:5555/api/v1/user/connect/${uid}/g/${gid}`;
    return await fetch(url);
}

export let connectDevice = async (uid, did) => {
    let url = `http://localhost:5555/api/v1/user/connect/${uid}/d/${did}`;
    return await fetch(url);
}

export let disconnectGroup = async (uid, gid) => {
    let url = `http://localhost:5555/api/v1/user/disconnect/${uid}/g/${gid}`;
    return await fetch(url);
}

export let disconnectDevice = async (uid, did) => {
    let url = `http://localhost:5555/api/v1/user/disconnect/${uid}/d/${did}`;
    return await fetch(url);
}

export let getUser = async (username) => {
    let url = `http://127.0.0.1:5555/api/v1/user/get/u/${username}`;
        
    let response = await fetch(url);
    console.log(response);
    let data = await response.json();
    console.log(data);
    return data;
}

export let getAllUser = async () => {
    let response = await fetch("http://127.0.0.1:5555/api/v1/user/get/all");
    let commits = await response.json();
    return commits;
}