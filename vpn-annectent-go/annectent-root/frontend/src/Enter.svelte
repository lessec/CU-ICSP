<script>
    // amazon : username: "loginwithamazon_amzn1.account.aguk7edgwdw55ttwfxwivifk2iwq"
    // google : username: "google_118202375641812213424"
    // normal
    // attributes: {
    //   custom:usertype: "Admin"
    //   email: "dudnspa0203@naver.com"
    //   email_verified: true
    //   sub: "9136ca50-e59c-4e58-bfc1-30ac24d892d1"
    // }
    // id: "ap-southeast-1:fb3b6699-52ed-4b7e-b387-324195ae432f"
    // username: "test0000"

    // amazon
    // attributes: {
    //   email: "nv.noelvalent@gmail.com"
    //   email_verified: false
    //   identities: "[{\"userId\":\"amzn1.account.AGUK7EDGWDW55TTWFXWIVIFK2IWQ\",\"providerName\":\"LoginWithAmazon\",\"providerType\":\"LoginWithAmazon\",\"issuer\":null,\"pri…"
    //   name: "Kim DongGyu"
    //   sub: "b35ed92e-e532-405b-9ed7-e3d4fe35af84"
    // }
    // id: "ap-southeast-1:ffd288cb-5766-4632-abf2-e6ed20af984f"
    // username: "loginwithamazon_amzn1.account.aguk7edgwdw55ttwfxwivifk2iwq"

    // google
    // attributes: {
    //   email: "nv.noelvalent@gmail.com"
    //   email_verified: false
    //   identities: "[{\"userId\":\"118202375641812213424\",\"providerName\":\"Google\",\"providerType\":\"Google\",\"issuer\":null,\"primary\":true,\"dateCreated\":1645089685759}…"
    //   sub: "ed15c647-76a2-45b4-aaae-9e55d86de363"
    // }
    // id: "ap-southeast-1:b6177938-1e87-40ab-a732-8199cfcf2d5a"
    // username: "google_118202375641812213424"
    import { Auth } from 'aws-amplify';
    import { onMount } from "svelte";
    import {UserStore} from "./Stores/user-store.js";
    import { navigate } from "svelte-routing";

    let isValid;
    let user;
	
    onMount(async () => {
        console.log($UserStore);
    if ($UserStore==null) {
            user = await Auth.currentAuthenticatedUser()
            .then((data)=>{return data})
            .catch((err)=>{return null});
            console.log(user)
            if (user==null) {
        // when any account is not logged in, redirect to login page
                navigate("/login", {replace: true});
            } else {
                // when any account is existsted in congnito but not in db.
                console.log(user);
                if (user.username.startsWith('loginwithamazon')
                ||user.username.startsWith('google_')) {
                    isValid = await isAccountValid(
                    user.attributes.name,
                    user.attributes.email,
                    user.attributes["custom:usertype"],
                    )
                } else {
                    isValid = await isAccountValid(
                    user.username,
                    user.attributes.email,
                    user.attributes["custom:usertype"],
                    )
                }
            
                if (isValid == false) {
                    navigate("/exit", {replace: true});
                } else {
                    // there is a valid account in db, but store is null;
                    $UserStore = user;
                    // goto dashboard
                    goToDashboard(); 
                }
            }
        } else {
        // goto dashboard
            goToDashboard();
        }
    });

    function goToDashboard() {
        // navigate("/usergroup", {replace: true});
        console.log("user store: ", $UserStore);
        console.log("user: ", user);
        if ($UserStore.attributes["custom:usertype"] == "Staff") {
            navigate("/admin", {replace: true});
        } else {
            navigate("/admin", {replace: true});
        }
    }

    async function isAccountValid(name, email, type) {
        let url = 'http://localhost:5555/api/v1/user/validation?name='+name+'&email='+email+'&type='+type;
        let response = await fetch(url).then(data=>{return data.json()}); 
        console.log(response)
		
        if (type==undefined || type==null) {
            let result = Auth.updateUserAttributes(user, {
                'custom:usertype': 'Staff',
            });
        }
        return response.isVaild;
    }
</script>