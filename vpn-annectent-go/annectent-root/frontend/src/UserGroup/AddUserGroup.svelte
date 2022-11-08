<script>
    /*
        This is Add User Group Component
    */
    import Button, { Label } from '@smui/button';
    import Textfield from '@smui/textfield';
    import CommonAppBar from '../Components/CommonAppBar.svelte';

    import { navigate } from "svelte-routing";

    let name = ''
    let describe = ''
    let grantedServer = ''

    // request add a group
    async function addUserGroup() {
        let url = 'http://localhost:5555/api/v1/group/add';

        const params = {
            name,
            describe,
            grantedServer,
        };

        const options =  {
            mode: "no-cors",
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
            body: JSON.stringify(params)
        }

        // TODO: how to check http status code? currently it is always 0.
        let res = await fetch(url, options).then(res => {
            console.log("new group added!");
            alert("Success!");
            navigate("/usergroup", {replace: true});
        });
    }

    function goToBack() {
        navigate("/usergroup", {replace: true});
    }
</script>

<CommonAppBar />

<main>
    <div class="inputs">
        <div class="lable-input">
            <div class="mdc-typography--headline4">Name</div>
            <Textfield bind:value={name} label="Name" style="width: 70%" />
        </div>
        <div class="lable-input">
            <div class="mdc-typography--headline4">Describe</div>
            <Textfield bind:value={describe} label="Describe" style="width: 70%" />
        </div>
    </div>
    <div class="buttons">
        <Button on:click={goToBack} variant="outlined" style="min-width: 170px">
            <Label>Cancel</Label>
        </Button>
        <Button on:click={addUserGroup} variant="unelevated" style="min-width: 170px">
            <Label>Save</Label>
        </Button>
    </div>
</main>

<style>
    main {
        display: flex;
        flex-direction: column;
        justify-content: center;
        margin: 5%;
    }
    main > div.inputs {
        display: flex;
        flex-direction: column;
    }
    main > div.buttons {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        margin-top: 18%;
    }
    
    div.inputs > div.lable-input {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
    }
</style>
