<script>
    import Textfield from '@smui/textfield';
    import HelperText from '@smui/textfield/helper-text';
    import Button, { Label } from '@smui/button';

    import { Auth } from 'aws-amplify';
    import { navigate } from "svelte-routing";
    import { UserStore } from "../Stores/user-store.js";
    import { addUser } from "../User/common-utils.js";
    import Logo from '../Assets/Logo.svelte';

    let code = '';
    let user = $UserStore;
    
    async function confirmSignUp() {
        try {
            await Auth.confirmSignUp(user.name, code).then(async () => {
                alert("Successfully confirmed sign up!");
                await Auth.signIn(user.name, user.password).then(async ()=>{
                    // add user to database
                    await addUser(user.name, user.email, user.type).then(()=>{
                        navigate("/", {replace: true});
                    });
                });
            });
            
        } catch (error) {
            console.log('error confirming sign up', error);
        }
    }

    async function resendConfirmationCode() {
        try {
            await Auth.resendSignUp(user.name);
            console.log('code resent successfully');
        } catch (err) {
            console.log('error resending code: ', err);
        }
    }

</script>

<main>
    <div class="inner">
        <Logo />
        <div class="mdc-typography--headline2 mt">Almost Done!</div>
        <div class="mdc-typography--subtitle1">You have to confirm your account.</div>
        <Textfield variant="outlined" bind:value={code} label="Code" style="min-width: 300px; margin-top: 30px">
            <HelperText slot="helper">Enter the Code</HelperText>
        </Textfield>
        <Button color="secondary" on:click={confirmSignUp} variant="unelevated" style="min-width: 300px; margin-top: 10px">
            <Label>Check my code!</Label>
        </Button>
        <Button color="secondary" on:click={resendConfirmationCode} variant="unelevated" style="min-width: 300px; margin-top: 10px">
            <Label>Resend mail</Label>
        </Button>
    </div>
</main>

<style>
    main {
        display: flex;
        justify-content: center;
    }
    div.inner {
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    .mt {
        margin-top: 40px;
    }
</style>