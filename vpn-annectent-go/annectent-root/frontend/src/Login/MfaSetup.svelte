<script>
	import Textfield from '@smui/textfield';
    import HelperText from '@smui/textfield/helper-text';
    import Button, { Label } from '@smui/button';

    import { Auth } from 'aws-amplify';
    import QrCode from 'svelte-qrcode';
    import {UserStore} from '../Stores/user-store.js';
    import { onMount } from 'svelte';
	import { navigate } from "svelte-routing";
	import CommonAppBar from '../Components/CommonAppBar.svelte';

    let user = $UserStore;
    let qrdata = "";
	let verifiedcode = "";

	console.log(user);
    function startTOPT() {
		// To setup TOTP, first you need to get a `authorization code` from Amazon Cognito
		// `user` is the current Authenticated user
		Auth.setupTOTP(user).then((code) => {
			// You can directly display the `code` to the user or convert it to a QR code to be scanned.
			// E.g., use following code sample to render a QR code with `qrcode.react` component:  
			//      import QRCode from 'qrcode.react';
			//      const str = "otpauth://totp/AWSCognito:"+ username + "?secret=" + code + "&issuer=" + issuer;
			//      <QRCode value={str}/>
			let issuer = "annectent";
			console.log(code);
			qrdata = "otpauth://totp/AWSCognito:"+ user.username + "?secret=" + code + "&issuer=" + issuer;
		});
	}

    onMount(() => {
        startTOPT();
    });

    function valifyTOTP() {
		// Then you will have your TOTP account in your TOTP-generating app (like Google Authenticator)
		// Use the generated one-time password to verify the setup

		let challengeAnswer = verifiedcode;
		Auth.verifyTotpToken(user, challengeAnswer).then(() => {
			// don't forget to set TOTP as the preferred MFA method
			Auth.setPreferredMFA(user, 'TOTP');
			navigate('/', {replace: true});
			}).catch( e => {
			// Token is not verified
			console.log(e);
		});
	}
</script>

<CommonAppBar />

<main>
    <div class="inner">
        <div class="mdc-typography--headline2 mt">This is MFA Setup Page</div>
        <div class="mdc-typography--subtitle1">You can regist MFA.</div>
		{#if qrdata != ""}
		<div class="container">
			<QrCode value={qrdata} />
		</div>
		  <Textfield variant="outlined" bind:value={verifiedcode} label="Verified Code" style="min-width: 300px">
			<HelperText slot="helper">Enter the Code</HelperText>
		  </Textfield>
		  <Button color="secondary" on:click={valifyTOTP} variant="unelevated" style="min-width: 300px; margin-top: 10px">
            <Label>Verify</Label>
        </Button>
		{/if}
        
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
