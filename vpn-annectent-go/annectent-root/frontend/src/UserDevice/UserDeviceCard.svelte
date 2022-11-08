<script>
    import { navigate } from "svelte-routing";

    export let device;
    export let getAlls;
    
    let id = device.id;
    let imgUrl = `http://127.0.0.1:5555/api/v1/device/get/id/${id}?format=qr&a=${Math.random()}`;

    // request to delete a device by id.
    async function deleteUserDevice() {
        let url = `http://127.0.0.1:5555/api/v1/device/delete/${id}`;

        const options = {
            method: 'POST',
        }

        await fetch(url, options).then(res => {
            console.log(res);
            alert("Device deleted");
            getAlls();
        });
    }

    // go to edit device page
    function editUserDevice() {
        navigate("device/edit/" + id, {replace: true});
    }
</script>

<div>
	<h3>{device.name}</h3>
    <ul>
        <li>IP = {device.ip}</li>
        <li>PublicKey = {device.publicKey}</li>
        <li>QRCode = {device.qr_code}</li>
        <li>ConfigFile = {device.configFile}</li>
    </ul>
    <img src="{imgUrl}" class="qrcode float-right" alt="Mobile client config"/>
    <button on:click={deleteUserDevice}>x</button>
    <button on:click={editUserDevice}>edit</button>
</div>