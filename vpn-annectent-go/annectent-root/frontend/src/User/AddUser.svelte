<script>
    import Button, { Label } from '@smui/button';
    import Textfield from '@smui/textfield';
    import FormField from '@smui/form-field';
    import Radio from '@smui/radio';
    import CommonAppBar from '../Components/CommonAppBar.svelte';
    
    import {navigate} from "svelte-routing";

    let name = '';
    let email = '';
    let type = '';
    let group = '';

    // request to add an user
    async function addUser() {
        let url = 'http://localhost:5555/api/v1/user/add';

        const params = {
            name,
            email,
            type,
            group,
        };

        const options = {
            mode: "no-cors",
            method: "POST",
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
            body: JSON.stringify(params)
        }

        await fetch(url, options).then(res => {
            console.log(res);
            alert("Success");
            navigate("/user", {replace: true})
        });

    }

    function goToBack() {
        navigate("/user", {replace: true})
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
            <div class="mdc-typography--headline4">User type</div>
            {#each ['Admin', 'Manager', 'Staff'] as option}
            <FormField style="margin-right: 1em;">
                <Radio bind:group={type} value={option} />
                <span slot="label">{`${option[0].toUpperCase()}${option.slice(1)}`}</span>
            </FormField>
            {/each}
        </div>
    </div>
    <div class="buttons">
        <Button on:click={goToBack} variant="outlined" style="min-width: 170px">
            <Label>Cancel</Label>
        </Button>
        <Button on:click={addUser} variant="unelevated" style="min-width: 170px">
            <Label>Save</Label>
        </Button>
    </div>
</main>

<style>
    main {
        display: flex;
        flex-direction: column;
        justify-content: center;
        margin: 10%;
        margin-top: 5%;
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
        margin-bottom: 30px;
    }
</style>