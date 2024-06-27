# WlanPortalTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessCodeAlternateEmail** | Pointer to **string** | “Please provide valid alternate email” | [optional] 
**Alignment** | Pointer to **string** | defines alignment on portal. “left” is default. | [optional] 
**AuthButtonAmazon** | Pointer to **string** | label for Amazon auth button | [optional] 
**AuthButtonAzure** | Pointer to **string** | label for Azure auth button | [optional] 
**AuthButtonEmail** | Pointer to **string** | label for Email auth button | [optional] 
**AuthButtonFacebook** | Pointer to **string** | label for Facebook auth button | [optional] 
**AuthButtonGoogle** | Pointer to **string** | label for Google auth button | [optional] 
**AuthButtonMicrosoft** | Pointer to **string** | label for Microsoft auth button | [optional] 
**AuthButtonPassphrase** | Pointer to **string** | label for passphrase auth button | [optional] 
**AuthButtonSms** | Pointer to **string** | label for SMS auth button | [optional] 
**AuthButtonSponsor** | Pointer to **string** | label for Sponsor auth button | [optional] 
**AuthLabel** | Pointer to **string** | “Connect to WiFi with” | [optional] 
**BackLink** | Pointer to **string** | label of the link to go back to /logon | [optional] 
**Color** | Pointer to **string** | “#1074bc” | [optional] 
**ColorDark** | Pointer to **string** | “#0b5183” | [optional] 
**ColorLight** | Pointer to **string** | “#3589c6” | [optional] 
**Company** | Pointer to **bool** | whether company field is required | [optional] [default to false]
**CompanyError** | Pointer to **string** | error message when company not provided | [optional] 
**CompanyLabel** | Pointer to **string** | label of company field | [optional] 
**Email** | Pointer to **bool** | whether email field is required | [optional] [default to false]
**EmailAccessDomainError** | Pointer to **string** | error message when a user has valid social login but doesn’t match specified email domains. | [optional] 
**EmailCancel** | Pointer to **string** | Label for cancel confirmation code submission using email auth | [optional] 
**EmailCodeCancel** | Pointer to **string** |  | [optional] 
**EmailCodeError** | Pointer to **string** | “Please provide valid alternate email” | [optional] 
**EmailCodeFieldLabel** | Pointer to **string** | “Confirmation Code” | [optional] 
**EmailCodeMessage** | Pointer to **string** | “Enter the access number that was sent to your email address.” | [optional] 
**EmailCodeSubmit** | Pointer to **string** | “Sign In | [optional] 
**EmailCodeTitle** | Pointer to **string** | “Access Code” | [optional] 
**EmailError** | Pointer to **string** | error message when email not provided | [optional] 
**EmailFieldLabel** | Pointer to **string** | “Enter your email address” | [optional] 
**EmailLabel** | Pointer to **string** | label of email field | [optional] 
**EmailMessage** | Pointer to **string** | “We will email you an authentication code which you can use to connect to the WiFi network.” | [optional] 
**EmailSubmit** | Pointer to **string** | Label for confirmation code submit button using email auth | [optional] 
**EmailTitle** | Pointer to **string** | “Sign in with Email” | [optional] 
**Field1** | Pointer to **bool** | whether to ask field1 | [optional] 
**Field1Error** | Pointer to **string** | error message when field1 not provided | [optional] 
**Field1Label** | Pointer to **string** | label of field1 | [optional] 
**Field1Required** | Pointer to **bool** | whether field1 is required field | [optional] 
**Field2** | Pointer to **bool** | whether to ask field2 | [optional] 
**Field2Error** | Pointer to **string** | error message when field2 not provided | [optional] 
**Field2Label** | Pointer to **string** | label of field2 | [optional] 
**Field2Required** | Pointer to **bool** | whether field2 is required field | [optional] 
**Field3** | Pointer to **bool** | whether to ask field3 | [optional] 
**Field3Error** | Pointer to **string** | error message when field3 not provided | [optional] 
**Field3Label** | Pointer to **string** | label of field3 | [optional] 
**Field3Required** | Pointer to **bool** | whether field3 is required field | [optional] 
**Field4** | Pointer to **bool** | whether to ask field4 | [optional] 
**Field4Error** | Pointer to **string** | error message when field4 not provided | [optional] 
**Field4Label** | Pointer to **string** | label of field4 | [optional] 
**Field4Required** | Pointer to **bool** | whether field4 is required field | [optional] 
**Message** | Pointer to **string** | “Please enjoy the complimentary Wifi” | [optional] 
**Name** | Pointer to **bool** | whether name field is required | [optional] [default to false]
**NameError** | Pointer to **string** | error message when name not provided | [optional] 
**NameLabel** | Pointer to **string** | label of name field | [optional] 
**Optout** | Pointer to **bool** | whether to display “Do Not Store My Personal Information” | [optional] [default to false]
**OptoutDefault** | Pointer to **bool** | Default value for the \\\&quot;Do not store\\\&quot; checkbox\&quot; | [optional] [default to true]
**OptoutLabel** | Pointer to **string** | label for “Do Not Store My Personal Information” | [optional] [default to "Do not store"]
**PageTitle** | **string** | “Welcome” | 
**PassphraseCancel** | Pointer to **string** | “Cancel” | [optional] 
**PassphraseError** | Pointer to **string** | error message when invalid passphrase is provided | [optional] 
**PassphraseLabel** | Pointer to **string** | Passphrase | [optional] 
**PassphraseMessage** | Pointer to **string** | “Login using passphrase” | [optional] 
**PassphraseSubmit** | Pointer to **string** | “Sign in” | [optional] 
**PassphraseTitle** | Pointer to **string** | Title for passphrase details page | [optional] 
**PoweredBy** | Pointer to **bool** | whether to show “Powered by Mist” | [optional] [default to true]
**RequiredFieldLabel** | Pointer to **string** | label to denote required field | [optional] 
**SignInLabel** | Pointer to **string** | label of the button to /signin | [optional] 
**SmsCarrierDefault** | Pointer to **string** | “Please Select” | [optional] 
**SmsCarrierError** | Pointer to **string** | “Please select a mobile carrier” | [optional] 
**SmsCarrierFieldLabel** | Pointer to **string** | label for mobile carrier drop-down list | [optional] 
**SmsCodeCancel** | Pointer to **string** | Label for cancel confirmation code submission | [optional] 
**SmsCodeError** | Pointer to **string** | error message when confirmation code is invalid | [optional] 
**SmsCodeFieldLabel** | Pointer to **string** | “Confirmation Code” | [optional] 
**SmsCodeMessage** | Pointer to **string** | “Enter the confirmation code” | [optional] 
**SmsCodeSubmit** | Pointer to **string** | Label for confirmation code submit button | [optional] 
**SmsCodeTitle** | Pointer to **string** | “Access Code” | [optional] 
**SmsCountryFieldLabel** | Pointer to **string** | “Country Code” | [optional] 
**SmsCountryFormat** | Pointer to **string** | “+1” | [optional] 
**SmsHaveAccessCode** | Pointer to **string** | Label for checkbox to specify that the user has access code | [optional] 
**SmsMessageFormat** | Pointer to **string** | format of access code sms message. {{code}} and {{duration}} are place holders and should be retained as is. | [optional] 
**SmsNumberCancel** | Pointer to **string** | label for canceling mobile details for SMS auth | [optional] 
**SmsNumberError** | Pointer to **string** | “Invalid Mobile Number” | [optional] 
**SmsNumberFieldLabel** | Pointer to **string** | label for field to provide mobile number | [optional] 
**SmsNumberFormat** | Pointer to **string** | “2125551212 (digits only)” | [optional] 
**SmsNumberMessage** | Pointer to **string** | “We will send an access code to your mobile number which you can use to connect to the WiFi network. Message and data rates may apply.” | [optional] 
**SmsNumberSubmit** | Pointer to **string** | label for submit button for code generation | [optional] 
**SmsNumberTitle** | Pointer to **string** | Title for phone number details | [optional] 
**SmsUsernameFormat** | Pointer to **string** | “username” | [optional] 
**SmsValidityDuration** | Pointer to **int32** | how long confirmation code should be considered valid (in minutes) | [optional] 
**SponsorBackLink** | Pointer to **string** | “Go back and edit request form” | [optional] 
**SponsorCancel** | Pointer to **string** | “Cancel” | [optional] 
**SponsorEmail** | Pointer to **string** | label for Sponsor Email | [optional] 
**SponsorEmailError** | Pointer to **string** | “Please provide valid sponsor email” | [optional] 
**SponsorEmailTemplate** | Pointer to **string** | html template to replace/override default sponsor email template  Sponsor Email Template supports following template variables:   * &#x60;approve_url&#x60;: Renders URL to approve the request; optionally &amp;minutes&#x3D;N query param can be appended to change the Authorization period of the guest, where N is a valid integer denoting number of minutes a guest remains authorized  * &#x60;deny_url&#x60;: Renders URL to reject the request  * &#x60;guest_email&#x60;: Renders Email ID of the guest  * &#x60;guest_name&#x60;: Renders Name of the guest  * &#x60;field1&#x60;: Renders value of the Custom Field 1  * &#x60;field2&#x60;: Renders value of the Custom Field 2  * &#x60;sponsor_link_validity_duration&#x60;: Renders validity time of the request (i.e. Approve/Deny URL)  * &#x60;auth_expire_minutes&#x60;: Renders Wlan-level configured Guest Authorization Expiration time period (in minutes), If not configured then default (1 day in minutes)  | [optional] 
**SponsorInfoApproved** | Pointer to **string** | “Your request was approved by” | [optional] 
**SponsorInfoDenied** | Pointer to **string** | “Your request was denied by” | [optional] 
**SponsorInfoPending** | Pointer to **string** | “Your notification has been sent to” | [optional] 
**SponsorName** | Pointer to **string** | label for Sponsor Name | [optional] 
**SponsorNameError** | Pointer to **string** | “Please provide sponsor’s name” | [optional] 
**SponsorNotePending** | Pointer to **string** | “Please wait for them to acknowledge.” | [optional] 
**SponsorRequestAccess** | Pointer to **string** | ‘submit button label request Wifi Access and notify sponsor about guest request | [optional] 
**SponsorSelectEmail** | Pointer to **string** | “Select Sponsor” | [optional] 
**SponsorStatusApproved** | Pointer to **string** | text to display if sponsor approves request | [optional] 
**SponsorStatusDenied** | Pointer to **string** | text to display when sponsor denies request | [optional] 
**SponsorStatusPending** | Pointer to **string** | text to display if request is still pending | [optional] 
**SponsorSubmit** | Pointer to **string** | submit button label to notify sponsor about guest request | [optional] 
**SponsorsError** | Pointer to **string** | “Please select a sponsor” | [optional] 
**SponsorsInfoApproved** | Pointer to **string** | “Your request was approved” | [optional] 
**SponsorsInfoDenied** | Pointer to **string** | “Your request was denied” | [optional] 
**SponsorsInfoPending** | Pointer to **string** | “Your notification has been sent to the sponsors” | [optional] 
**Tos** | Pointer to **bool** |  | [optional] [default to true]
**TosAcceptLabel** | Pointer to **string** | prefix of the label of the link to go to /tos | [optional] 
**TosError** | Pointer to **string** | error message when tos not accepted | [optional] 
**TosLink** | Pointer to **string** | label of the link to go to /tos | [optional] 
**TosText** | Pointer to **string** | text of the Terms of Service | [optional] 

## Methods

### NewWlanPortalTemplate

`func NewWlanPortalTemplate(pageTitle string, ) *WlanPortalTemplate`

NewWlanPortalTemplate instantiates a new WlanPortalTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanPortalTemplateWithDefaults

`func NewWlanPortalTemplateWithDefaults() *WlanPortalTemplate`

NewWlanPortalTemplateWithDefaults instantiates a new WlanPortalTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessCodeAlternateEmail

`func (o *WlanPortalTemplate) GetAccessCodeAlternateEmail() string`

GetAccessCodeAlternateEmail returns the AccessCodeAlternateEmail field if non-nil, zero value otherwise.

### GetAccessCodeAlternateEmailOk

`func (o *WlanPortalTemplate) GetAccessCodeAlternateEmailOk() (*string, bool)`

GetAccessCodeAlternateEmailOk returns a tuple with the AccessCodeAlternateEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCodeAlternateEmail

`func (o *WlanPortalTemplate) SetAccessCodeAlternateEmail(v string)`

SetAccessCodeAlternateEmail sets AccessCodeAlternateEmail field to given value.

### HasAccessCodeAlternateEmail

`func (o *WlanPortalTemplate) HasAccessCodeAlternateEmail() bool`

HasAccessCodeAlternateEmail returns a boolean if a field has been set.

### GetAlignment

`func (o *WlanPortalTemplate) GetAlignment() string`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *WlanPortalTemplate) GetAlignmentOk() (*string, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *WlanPortalTemplate) SetAlignment(v string)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *WlanPortalTemplate) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetAuthButtonAmazon

`func (o *WlanPortalTemplate) GetAuthButtonAmazon() string`

GetAuthButtonAmazon returns the AuthButtonAmazon field if non-nil, zero value otherwise.

### GetAuthButtonAmazonOk

`func (o *WlanPortalTemplate) GetAuthButtonAmazonOk() (*string, bool)`

GetAuthButtonAmazonOk returns a tuple with the AuthButtonAmazon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonAmazon

`func (o *WlanPortalTemplate) SetAuthButtonAmazon(v string)`

SetAuthButtonAmazon sets AuthButtonAmazon field to given value.

### HasAuthButtonAmazon

`func (o *WlanPortalTemplate) HasAuthButtonAmazon() bool`

HasAuthButtonAmazon returns a boolean if a field has been set.

### GetAuthButtonAzure

`func (o *WlanPortalTemplate) GetAuthButtonAzure() string`

GetAuthButtonAzure returns the AuthButtonAzure field if non-nil, zero value otherwise.

### GetAuthButtonAzureOk

`func (o *WlanPortalTemplate) GetAuthButtonAzureOk() (*string, bool)`

GetAuthButtonAzureOk returns a tuple with the AuthButtonAzure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonAzure

`func (o *WlanPortalTemplate) SetAuthButtonAzure(v string)`

SetAuthButtonAzure sets AuthButtonAzure field to given value.

### HasAuthButtonAzure

`func (o *WlanPortalTemplate) HasAuthButtonAzure() bool`

HasAuthButtonAzure returns a boolean if a field has been set.

### GetAuthButtonEmail

`func (o *WlanPortalTemplate) GetAuthButtonEmail() string`

GetAuthButtonEmail returns the AuthButtonEmail field if non-nil, zero value otherwise.

### GetAuthButtonEmailOk

`func (o *WlanPortalTemplate) GetAuthButtonEmailOk() (*string, bool)`

GetAuthButtonEmailOk returns a tuple with the AuthButtonEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonEmail

`func (o *WlanPortalTemplate) SetAuthButtonEmail(v string)`

SetAuthButtonEmail sets AuthButtonEmail field to given value.

### HasAuthButtonEmail

`func (o *WlanPortalTemplate) HasAuthButtonEmail() bool`

HasAuthButtonEmail returns a boolean if a field has been set.

### GetAuthButtonFacebook

`func (o *WlanPortalTemplate) GetAuthButtonFacebook() string`

GetAuthButtonFacebook returns the AuthButtonFacebook field if non-nil, zero value otherwise.

### GetAuthButtonFacebookOk

`func (o *WlanPortalTemplate) GetAuthButtonFacebookOk() (*string, bool)`

GetAuthButtonFacebookOk returns a tuple with the AuthButtonFacebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonFacebook

`func (o *WlanPortalTemplate) SetAuthButtonFacebook(v string)`

SetAuthButtonFacebook sets AuthButtonFacebook field to given value.

### HasAuthButtonFacebook

`func (o *WlanPortalTemplate) HasAuthButtonFacebook() bool`

HasAuthButtonFacebook returns a boolean if a field has been set.

### GetAuthButtonGoogle

`func (o *WlanPortalTemplate) GetAuthButtonGoogle() string`

GetAuthButtonGoogle returns the AuthButtonGoogle field if non-nil, zero value otherwise.

### GetAuthButtonGoogleOk

`func (o *WlanPortalTemplate) GetAuthButtonGoogleOk() (*string, bool)`

GetAuthButtonGoogleOk returns a tuple with the AuthButtonGoogle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonGoogle

`func (o *WlanPortalTemplate) SetAuthButtonGoogle(v string)`

SetAuthButtonGoogle sets AuthButtonGoogle field to given value.

### HasAuthButtonGoogle

`func (o *WlanPortalTemplate) HasAuthButtonGoogle() bool`

HasAuthButtonGoogle returns a boolean if a field has been set.

### GetAuthButtonMicrosoft

`func (o *WlanPortalTemplate) GetAuthButtonMicrosoft() string`

GetAuthButtonMicrosoft returns the AuthButtonMicrosoft field if non-nil, zero value otherwise.

### GetAuthButtonMicrosoftOk

`func (o *WlanPortalTemplate) GetAuthButtonMicrosoftOk() (*string, bool)`

GetAuthButtonMicrosoftOk returns a tuple with the AuthButtonMicrosoft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonMicrosoft

`func (o *WlanPortalTemplate) SetAuthButtonMicrosoft(v string)`

SetAuthButtonMicrosoft sets AuthButtonMicrosoft field to given value.

### HasAuthButtonMicrosoft

`func (o *WlanPortalTemplate) HasAuthButtonMicrosoft() bool`

HasAuthButtonMicrosoft returns a boolean if a field has been set.

### GetAuthButtonPassphrase

`func (o *WlanPortalTemplate) GetAuthButtonPassphrase() string`

GetAuthButtonPassphrase returns the AuthButtonPassphrase field if non-nil, zero value otherwise.

### GetAuthButtonPassphraseOk

`func (o *WlanPortalTemplate) GetAuthButtonPassphraseOk() (*string, bool)`

GetAuthButtonPassphraseOk returns a tuple with the AuthButtonPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonPassphrase

`func (o *WlanPortalTemplate) SetAuthButtonPassphrase(v string)`

SetAuthButtonPassphrase sets AuthButtonPassphrase field to given value.

### HasAuthButtonPassphrase

`func (o *WlanPortalTemplate) HasAuthButtonPassphrase() bool`

HasAuthButtonPassphrase returns a boolean if a field has been set.

### GetAuthButtonSms

`func (o *WlanPortalTemplate) GetAuthButtonSms() string`

GetAuthButtonSms returns the AuthButtonSms field if non-nil, zero value otherwise.

### GetAuthButtonSmsOk

`func (o *WlanPortalTemplate) GetAuthButtonSmsOk() (*string, bool)`

GetAuthButtonSmsOk returns a tuple with the AuthButtonSms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonSms

`func (o *WlanPortalTemplate) SetAuthButtonSms(v string)`

SetAuthButtonSms sets AuthButtonSms field to given value.

### HasAuthButtonSms

`func (o *WlanPortalTemplate) HasAuthButtonSms() bool`

HasAuthButtonSms returns a boolean if a field has been set.

### GetAuthButtonSponsor

`func (o *WlanPortalTemplate) GetAuthButtonSponsor() string`

GetAuthButtonSponsor returns the AuthButtonSponsor field if non-nil, zero value otherwise.

### GetAuthButtonSponsorOk

`func (o *WlanPortalTemplate) GetAuthButtonSponsorOk() (*string, bool)`

GetAuthButtonSponsorOk returns a tuple with the AuthButtonSponsor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthButtonSponsor

`func (o *WlanPortalTemplate) SetAuthButtonSponsor(v string)`

SetAuthButtonSponsor sets AuthButtonSponsor field to given value.

### HasAuthButtonSponsor

`func (o *WlanPortalTemplate) HasAuthButtonSponsor() bool`

HasAuthButtonSponsor returns a boolean if a field has been set.

### GetAuthLabel

`func (o *WlanPortalTemplate) GetAuthLabel() string`

GetAuthLabel returns the AuthLabel field if non-nil, zero value otherwise.

### GetAuthLabelOk

`func (o *WlanPortalTemplate) GetAuthLabelOk() (*string, bool)`

GetAuthLabelOk returns a tuple with the AuthLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthLabel

`func (o *WlanPortalTemplate) SetAuthLabel(v string)`

SetAuthLabel sets AuthLabel field to given value.

### HasAuthLabel

`func (o *WlanPortalTemplate) HasAuthLabel() bool`

HasAuthLabel returns a boolean if a field has been set.

### GetBackLink

`func (o *WlanPortalTemplate) GetBackLink() string`

GetBackLink returns the BackLink field if non-nil, zero value otherwise.

### GetBackLinkOk

`func (o *WlanPortalTemplate) GetBackLinkOk() (*string, bool)`

GetBackLinkOk returns a tuple with the BackLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackLink

`func (o *WlanPortalTemplate) SetBackLink(v string)`

SetBackLink sets BackLink field to given value.

### HasBackLink

`func (o *WlanPortalTemplate) HasBackLink() bool`

HasBackLink returns a boolean if a field has been set.

### GetColor

`func (o *WlanPortalTemplate) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WlanPortalTemplate) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WlanPortalTemplate) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *WlanPortalTemplate) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetColorDark

`func (o *WlanPortalTemplate) GetColorDark() string`

GetColorDark returns the ColorDark field if non-nil, zero value otherwise.

### GetColorDarkOk

`func (o *WlanPortalTemplate) GetColorDarkOk() (*string, bool)`

GetColorDarkOk returns a tuple with the ColorDark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDark

`func (o *WlanPortalTemplate) SetColorDark(v string)`

SetColorDark sets ColorDark field to given value.

### HasColorDark

`func (o *WlanPortalTemplate) HasColorDark() bool`

HasColorDark returns a boolean if a field has been set.

### GetColorLight

`func (o *WlanPortalTemplate) GetColorLight() string`

GetColorLight returns the ColorLight field if non-nil, zero value otherwise.

### GetColorLightOk

`func (o *WlanPortalTemplate) GetColorLightOk() (*string, bool)`

GetColorLightOk returns a tuple with the ColorLight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorLight

`func (o *WlanPortalTemplate) SetColorLight(v string)`

SetColorLight sets ColorLight field to given value.

### HasColorLight

`func (o *WlanPortalTemplate) HasColorLight() bool`

HasColorLight returns a boolean if a field has been set.

### GetCompany

`func (o *WlanPortalTemplate) GetCompany() bool`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *WlanPortalTemplate) GetCompanyOk() (*bool, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *WlanPortalTemplate) SetCompany(v bool)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *WlanPortalTemplate) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyError

`func (o *WlanPortalTemplate) GetCompanyError() string`

GetCompanyError returns the CompanyError field if non-nil, zero value otherwise.

### GetCompanyErrorOk

`func (o *WlanPortalTemplate) GetCompanyErrorOk() (*string, bool)`

GetCompanyErrorOk returns a tuple with the CompanyError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyError

`func (o *WlanPortalTemplate) SetCompanyError(v string)`

SetCompanyError sets CompanyError field to given value.

### HasCompanyError

`func (o *WlanPortalTemplate) HasCompanyError() bool`

HasCompanyError returns a boolean if a field has been set.

### GetCompanyLabel

`func (o *WlanPortalTemplate) GetCompanyLabel() string`

GetCompanyLabel returns the CompanyLabel field if non-nil, zero value otherwise.

### GetCompanyLabelOk

`func (o *WlanPortalTemplate) GetCompanyLabelOk() (*string, bool)`

GetCompanyLabelOk returns a tuple with the CompanyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLabel

`func (o *WlanPortalTemplate) SetCompanyLabel(v string)`

SetCompanyLabel sets CompanyLabel field to given value.

### HasCompanyLabel

`func (o *WlanPortalTemplate) HasCompanyLabel() bool`

HasCompanyLabel returns a boolean if a field has been set.

### GetEmail

`func (o *WlanPortalTemplate) GetEmail() bool`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *WlanPortalTemplate) GetEmailOk() (*bool, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *WlanPortalTemplate) SetEmail(v bool)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *WlanPortalTemplate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailAccessDomainError

`func (o *WlanPortalTemplate) GetEmailAccessDomainError() string`

GetEmailAccessDomainError returns the EmailAccessDomainError field if non-nil, zero value otherwise.

### GetEmailAccessDomainErrorOk

`func (o *WlanPortalTemplate) GetEmailAccessDomainErrorOk() (*string, bool)`

GetEmailAccessDomainErrorOk returns a tuple with the EmailAccessDomainError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAccessDomainError

`func (o *WlanPortalTemplate) SetEmailAccessDomainError(v string)`

SetEmailAccessDomainError sets EmailAccessDomainError field to given value.

### HasEmailAccessDomainError

`func (o *WlanPortalTemplate) HasEmailAccessDomainError() bool`

HasEmailAccessDomainError returns a boolean if a field has been set.

### GetEmailCancel

`func (o *WlanPortalTemplate) GetEmailCancel() string`

GetEmailCancel returns the EmailCancel field if non-nil, zero value otherwise.

### GetEmailCancelOk

`func (o *WlanPortalTemplate) GetEmailCancelOk() (*string, bool)`

GetEmailCancelOk returns a tuple with the EmailCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCancel

`func (o *WlanPortalTemplate) SetEmailCancel(v string)`

SetEmailCancel sets EmailCancel field to given value.

### HasEmailCancel

`func (o *WlanPortalTemplate) HasEmailCancel() bool`

HasEmailCancel returns a boolean if a field has been set.

### GetEmailCodeCancel

`func (o *WlanPortalTemplate) GetEmailCodeCancel() string`

GetEmailCodeCancel returns the EmailCodeCancel field if non-nil, zero value otherwise.

### GetEmailCodeCancelOk

`func (o *WlanPortalTemplate) GetEmailCodeCancelOk() (*string, bool)`

GetEmailCodeCancelOk returns a tuple with the EmailCodeCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCodeCancel

`func (o *WlanPortalTemplate) SetEmailCodeCancel(v string)`

SetEmailCodeCancel sets EmailCodeCancel field to given value.

### HasEmailCodeCancel

`func (o *WlanPortalTemplate) HasEmailCodeCancel() bool`

HasEmailCodeCancel returns a boolean if a field has been set.

### GetEmailCodeError

`func (o *WlanPortalTemplate) GetEmailCodeError() string`

GetEmailCodeError returns the EmailCodeError field if non-nil, zero value otherwise.

### GetEmailCodeErrorOk

`func (o *WlanPortalTemplate) GetEmailCodeErrorOk() (*string, bool)`

GetEmailCodeErrorOk returns a tuple with the EmailCodeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCodeError

`func (o *WlanPortalTemplate) SetEmailCodeError(v string)`

SetEmailCodeError sets EmailCodeError field to given value.

### HasEmailCodeError

`func (o *WlanPortalTemplate) HasEmailCodeError() bool`

HasEmailCodeError returns a boolean if a field has been set.

### GetEmailCodeFieldLabel

`func (o *WlanPortalTemplate) GetEmailCodeFieldLabel() string`

GetEmailCodeFieldLabel returns the EmailCodeFieldLabel field if non-nil, zero value otherwise.

### GetEmailCodeFieldLabelOk

`func (o *WlanPortalTemplate) GetEmailCodeFieldLabelOk() (*string, bool)`

GetEmailCodeFieldLabelOk returns a tuple with the EmailCodeFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCodeFieldLabel

`func (o *WlanPortalTemplate) SetEmailCodeFieldLabel(v string)`

SetEmailCodeFieldLabel sets EmailCodeFieldLabel field to given value.

### HasEmailCodeFieldLabel

`func (o *WlanPortalTemplate) HasEmailCodeFieldLabel() bool`

HasEmailCodeFieldLabel returns a boolean if a field has been set.

### GetEmailCodeMessage

`func (o *WlanPortalTemplate) GetEmailCodeMessage() string`

GetEmailCodeMessage returns the EmailCodeMessage field if non-nil, zero value otherwise.

### GetEmailCodeMessageOk

`func (o *WlanPortalTemplate) GetEmailCodeMessageOk() (*string, bool)`

GetEmailCodeMessageOk returns a tuple with the EmailCodeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCodeMessage

`func (o *WlanPortalTemplate) SetEmailCodeMessage(v string)`

SetEmailCodeMessage sets EmailCodeMessage field to given value.

### HasEmailCodeMessage

`func (o *WlanPortalTemplate) HasEmailCodeMessage() bool`

HasEmailCodeMessage returns a boolean if a field has been set.

### GetEmailCodeSubmit

`func (o *WlanPortalTemplate) GetEmailCodeSubmit() string`

GetEmailCodeSubmit returns the EmailCodeSubmit field if non-nil, zero value otherwise.

### GetEmailCodeSubmitOk

`func (o *WlanPortalTemplate) GetEmailCodeSubmitOk() (*string, bool)`

GetEmailCodeSubmitOk returns a tuple with the EmailCodeSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCodeSubmit

`func (o *WlanPortalTemplate) SetEmailCodeSubmit(v string)`

SetEmailCodeSubmit sets EmailCodeSubmit field to given value.

### HasEmailCodeSubmit

`func (o *WlanPortalTemplate) HasEmailCodeSubmit() bool`

HasEmailCodeSubmit returns a boolean if a field has been set.

### GetEmailCodeTitle

`func (o *WlanPortalTemplate) GetEmailCodeTitle() string`

GetEmailCodeTitle returns the EmailCodeTitle field if non-nil, zero value otherwise.

### GetEmailCodeTitleOk

`func (o *WlanPortalTemplate) GetEmailCodeTitleOk() (*string, bool)`

GetEmailCodeTitleOk returns a tuple with the EmailCodeTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCodeTitle

`func (o *WlanPortalTemplate) SetEmailCodeTitle(v string)`

SetEmailCodeTitle sets EmailCodeTitle field to given value.

### HasEmailCodeTitle

`func (o *WlanPortalTemplate) HasEmailCodeTitle() bool`

HasEmailCodeTitle returns a boolean if a field has been set.

### GetEmailError

`func (o *WlanPortalTemplate) GetEmailError() string`

GetEmailError returns the EmailError field if non-nil, zero value otherwise.

### GetEmailErrorOk

`func (o *WlanPortalTemplate) GetEmailErrorOk() (*string, bool)`

GetEmailErrorOk returns a tuple with the EmailError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailError

`func (o *WlanPortalTemplate) SetEmailError(v string)`

SetEmailError sets EmailError field to given value.

### HasEmailError

`func (o *WlanPortalTemplate) HasEmailError() bool`

HasEmailError returns a boolean if a field has been set.

### GetEmailFieldLabel

`func (o *WlanPortalTemplate) GetEmailFieldLabel() string`

GetEmailFieldLabel returns the EmailFieldLabel field if non-nil, zero value otherwise.

### GetEmailFieldLabelOk

`func (o *WlanPortalTemplate) GetEmailFieldLabelOk() (*string, bool)`

GetEmailFieldLabelOk returns a tuple with the EmailFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFieldLabel

`func (o *WlanPortalTemplate) SetEmailFieldLabel(v string)`

SetEmailFieldLabel sets EmailFieldLabel field to given value.

### HasEmailFieldLabel

`func (o *WlanPortalTemplate) HasEmailFieldLabel() bool`

HasEmailFieldLabel returns a boolean if a field has been set.

### GetEmailLabel

`func (o *WlanPortalTemplate) GetEmailLabel() string`

GetEmailLabel returns the EmailLabel field if non-nil, zero value otherwise.

### GetEmailLabelOk

`func (o *WlanPortalTemplate) GetEmailLabelOk() (*string, bool)`

GetEmailLabelOk returns a tuple with the EmailLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailLabel

`func (o *WlanPortalTemplate) SetEmailLabel(v string)`

SetEmailLabel sets EmailLabel field to given value.

### HasEmailLabel

`func (o *WlanPortalTemplate) HasEmailLabel() bool`

HasEmailLabel returns a boolean if a field has been set.

### GetEmailMessage

`func (o *WlanPortalTemplate) GetEmailMessage() string`

GetEmailMessage returns the EmailMessage field if non-nil, zero value otherwise.

### GetEmailMessageOk

`func (o *WlanPortalTemplate) GetEmailMessageOk() (*string, bool)`

GetEmailMessageOk returns a tuple with the EmailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMessage

`func (o *WlanPortalTemplate) SetEmailMessage(v string)`

SetEmailMessage sets EmailMessage field to given value.

### HasEmailMessage

`func (o *WlanPortalTemplate) HasEmailMessage() bool`

HasEmailMessage returns a boolean if a field has been set.

### GetEmailSubmit

`func (o *WlanPortalTemplate) GetEmailSubmit() string`

GetEmailSubmit returns the EmailSubmit field if non-nil, zero value otherwise.

### GetEmailSubmitOk

`func (o *WlanPortalTemplate) GetEmailSubmitOk() (*string, bool)`

GetEmailSubmitOk returns a tuple with the EmailSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubmit

`func (o *WlanPortalTemplate) SetEmailSubmit(v string)`

SetEmailSubmit sets EmailSubmit field to given value.

### HasEmailSubmit

`func (o *WlanPortalTemplate) HasEmailSubmit() bool`

HasEmailSubmit returns a boolean if a field has been set.

### GetEmailTitle

`func (o *WlanPortalTemplate) GetEmailTitle() string`

GetEmailTitle returns the EmailTitle field if non-nil, zero value otherwise.

### GetEmailTitleOk

`func (o *WlanPortalTemplate) GetEmailTitleOk() (*string, bool)`

GetEmailTitleOk returns a tuple with the EmailTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTitle

`func (o *WlanPortalTemplate) SetEmailTitle(v string)`

SetEmailTitle sets EmailTitle field to given value.

### HasEmailTitle

`func (o *WlanPortalTemplate) HasEmailTitle() bool`

HasEmailTitle returns a boolean if a field has been set.

### GetField1

`func (o *WlanPortalTemplate) GetField1() bool`

GetField1 returns the Field1 field if non-nil, zero value otherwise.

### GetField1Ok

`func (o *WlanPortalTemplate) GetField1Ok() (*bool, bool)`

GetField1Ok returns a tuple with the Field1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField1

`func (o *WlanPortalTemplate) SetField1(v bool)`

SetField1 sets Field1 field to given value.

### HasField1

`func (o *WlanPortalTemplate) HasField1() bool`

HasField1 returns a boolean if a field has been set.

### GetField1Error

`func (o *WlanPortalTemplate) GetField1Error() string`

GetField1Error returns the Field1Error field if non-nil, zero value otherwise.

### GetField1ErrorOk

`func (o *WlanPortalTemplate) GetField1ErrorOk() (*string, bool)`

GetField1ErrorOk returns a tuple with the Field1Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField1Error

`func (o *WlanPortalTemplate) SetField1Error(v string)`

SetField1Error sets Field1Error field to given value.

### HasField1Error

`func (o *WlanPortalTemplate) HasField1Error() bool`

HasField1Error returns a boolean if a field has been set.

### GetField1Label

`func (o *WlanPortalTemplate) GetField1Label() string`

GetField1Label returns the Field1Label field if non-nil, zero value otherwise.

### GetField1LabelOk

`func (o *WlanPortalTemplate) GetField1LabelOk() (*string, bool)`

GetField1LabelOk returns a tuple with the Field1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField1Label

`func (o *WlanPortalTemplate) SetField1Label(v string)`

SetField1Label sets Field1Label field to given value.

### HasField1Label

`func (o *WlanPortalTemplate) HasField1Label() bool`

HasField1Label returns a boolean if a field has been set.

### GetField1Required

`func (o *WlanPortalTemplate) GetField1Required() bool`

GetField1Required returns the Field1Required field if non-nil, zero value otherwise.

### GetField1RequiredOk

`func (o *WlanPortalTemplate) GetField1RequiredOk() (*bool, bool)`

GetField1RequiredOk returns a tuple with the Field1Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField1Required

`func (o *WlanPortalTemplate) SetField1Required(v bool)`

SetField1Required sets Field1Required field to given value.

### HasField1Required

`func (o *WlanPortalTemplate) HasField1Required() bool`

HasField1Required returns a boolean if a field has been set.

### GetField2

`func (o *WlanPortalTemplate) GetField2() bool`

GetField2 returns the Field2 field if non-nil, zero value otherwise.

### GetField2Ok

`func (o *WlanPortalTemplate) GetField2Ok() (*bool, bool)`

GetField2Ok returns a tuple with the Field2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField2

`func (o *WlanPortalTemplate) SetField2(v bool)`

SetField2 sets Field2 field to given value.

### HasField2

`func (o *WlanPortalTemplate) HasField2() bool`

HasField2 returns a boolean if a field has been set.

### GetField2Error

`func (o *WlanPortalTemplate) GetField2Error() string`

GetField2Error returns the Field2Error field if non-nil, zero value otherwise.

### GetField2ErrorOk

`func (o *WlanPortalTemplate) GetField2ErrorOk() (*string, bool)`

GetField2ErrorOk returns a tuple with the Field2Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField2Error

`func (o *WlanPortalTemplate) SetField2Error(v string)`

SetField2Error sets Field2Error field to given value.

### HasField2Error

`func (o *WlanPortalTemplate) HasField2Error() bool`

HasField2Error returns a boolean if a field has been set.

### GetField2Label

`func (o *WlanPortalTemplate) GetField2Label() string`

GetField2Label returns the Field2Label field if non-nil, zero value otherwise.

### GetField2LabelOk

`func (o *WlanPortalTemplate) GetField2LabelOk() (*string, bool)`

GetField2LabelOk returns a tuple with the Field2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField2Label

`func (o *WlanPortalTemplate) SetField2Label(v string)`

SetField2Label sets Field2Label field to given value.

### HasField2Label

`func (o *WlanPortalTemplate) HasField2Label() bool`

HasField2Label returns a boolean if a field has been set.

### GetField2Required

`func (o *WlanPortalTemplate) GetField2Required() bool`

GetField2Required returns the Field2Required field if non-nil, zero value otherwise.

### GetField2RequiredOk

`func (o *WlanPortalTemplate) GetField2RequiredOk() (*bool, bool)`

GetField2RequiredOk returns a tuple with the Field2Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField2Required

`func (o *WlanPortalTemplate) SetField2Required(v bool)`

SetField2Required sets Field2Required field to given value.

### HasField2Required

`func (o *WlanPortalTemplate) HasField2Required() bool`

HasField2Required returns a boolean if a field has been set.

### GetField3

`func (o *WlanPortalTemplate) GetField3() bool`

GetField3 returns the Field3 field if non-nil, zero value otherwise.

### GetField3Ok

`func (o *WlanPortalTemplate) GetField3Ok() (*bool, bool)`

GetField3Ok returns a tuple with the Field3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField3

`func (o *WlanPortalTemplate) SetField3(v bool)`

SetField3 sets Field3 field to given value.

### HasField3

`func (o *WlanPortalTemplate) HasField3() bool`

HasField3 returns a boolean if a field has been set.

### GetField3Error

`func (o *WlanPortalTemplate) GetField3Error() string`

GetField3Error returns the Field3Error field if non-nil, zero value otherwise.

### GetField3ErrorOk

`func (o *WlanPortalTemplate) GetField3ErrorOk() (*string, bool)`

GetField3ErrorOk returns a tuple with the Field3Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField3Error

`func (o *WlanPortalTemplate) SetField3Error(v string)`

SetField3Error sets Field3Error field to given value.

### HasField3Error

`func (o *WlanPortalTemplate) HasField3Error() bool`

HasField3Error returns a boolean if a field has been set.

### GetField3Label

`func (o *WlanPortalTemplate) GetField3Label() string`

GetField3Label returns the Field3Label field if non-nil, zero value otherwise.

### GetField3LabelOk

`func (o *WlanPortalTemplate) GetField3LabelOk() (*string, bool)`

GetField3LabelOk returns a tuple with the Field3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField3Label

`func (o *WlanPortalTemplate) SetField3Label(v string)`

SetField3Label sets Field3Label field to given value.

### HasField3Label

`func (o *WlanPortalTemplate) HasField3Label() bool`

HasField3Label returns a boolean if a field has been set.

### GetField3Required

`func (o *WlanPortalTemplate) GetField3Required() bool`

GetField3Required returns the Field3Required field if non-nil, zero value otherwise.

### GetField3RequiredOk

`func (o *WlanPortalTemplate) GetField3RequiredOk() (*bool, bool)`

GetField3RequiredOk returns a tuple with the Field3Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField3Required

`func (o *WlanPortalTemplate) SetField3Required(v bool)`

SetField3Required sets Field3Required field to given value.

### HasField3Required

`func (o *WlanPortalTemplate) HasField3Required() bool`

HasField3Required returns a boolean if a field has been set.

### GetField4

`func (o *WlanPortalTemplate) GetField4() bool`

GetField4 returns the Field4 field if non-nil, zero value otherwise.

### GetField4Ok

`func (o *WlanPortalTemplate) GetField4Ok() (*bool, bool)`

GetField4Ok returns a tuple with the Field4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField4

`func (o *WlanPortalTemplate) SetField4(v bool)`

SetField4 sets Field4 field to given value.

### HasField4

`func (o *WlanPortalTemplate) HasField4() bool`

HasField4 returns a boolean if a field has been set.

### GetField4Error

`func (o *WlanPortalTemplate) GetField4Error() string`

GetField4Error returns the Field4Error field if non-nil, zero value otherwise.

### GetField4ErrorOk

`func (o *WlanPortalTemplate) GetField4ErrorOk() (*string, bool)`

GetField4ErrorOk returns a tuple with the Field4Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField4Error

`func (o *WlanPortalTemplate) SetField4Error(v string)`

SetField4Error sets Field4Error field to given value.

### HasField4Error

`func (o *WlanPortalTemplate) HasField4Error() bool`

HasField4Error returns a boolean if a field has been set.

### GetField4Label

`func (o *WlanPortalTemplate) GetField4Label() string`

GetField4Label returns the Field4Label field if non-nil, zero value otherwise.

### GetField4LabelOk

`func (o *WlanPortalTemplate) GetField4LabelOk() (*string, bool)`

GetField4LabelOk returns a tuple with the Field4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField4Label

`func (o *WlanPortalTemplate) SetField4Label(v string)`

SetField4Label sets Field4Label field to given value.

### HasField4Label

`func (o *WlanPortalTemplate) HasField4Label() bool`

HasField4Label returns a boolean if a field has been set.

### GetField4Required

`func (o *WlanPortalTemplate) GetField4Required() bool`

GetField4Required returns the Field4Required field if non-nil, zero value otherwise.

### GetField4RequiredOk

`func (o *WlanPortalTemplate) GetField4RequiredOk() (*bool, bool)`

GetField4RequiredOk returns a tuple with the Field4Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField4Required

`func (o *WlanPortalTemplate) SetField4Required(v bool)`

SetField4Required sets Field4Required field to given value.

### HasField4Required

`func (o *WlanPortalTemplate) HasField4Required() bool`

HasField4Required returns a boolean if a field has been set.

### GetMessage

`func (o *WlanPortalTemplate) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WlanPortalTemplate) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WlanPortalTemplate) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WlanPortalTemplate) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *WlanPortalTemplate) GetName() bool`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WlanPortalTemplate) GetNameOk() (*bool, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WlanPortalTemplate) SetName(v bool)`

SetName sets Name field to given value.

### HasName

`func (o *WlanPortalTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameError

`func (o *WlanPortalTemplate) GetNameError() string`

GetNameError returns the NameError field if non-nil, zero value otherwise.

### GetNameErrorOk

`func (o *WlanPortalTemplate) GetNameErrorOk() (*string, bool)`

GetNameErrorOk returns a tuple with the NameError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameError

`func (o *WlanPortalTemplate) SetNameError(v string)`

SetNameError sets NameError field to given value.

### HasNameError

`func (o *WlanPortalTemplate) HasNameError() bool`

HasNameError returns a boolean if a field has been set.

### GetNameLabel

`func (o *WlanPortalTemplate) GetNameLabel() string`

GetNameLabel returns the NameLabel field if non-nil, zero value otherwise.

### GetNameLabelOk

`func (o *WlanPortalTemplate) GetNameLabelOk() (*string, bool)`

GetNameLabelOk returns a tuple with the NameLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLabel

`func (o *WlanPortalTemplate) SetNameLabel(v string)`

SetNameLabel sets NameLabel field to given value.

### HasNameLabel

`func (o *WlanPortalTemplate) HasNameLabel() bool`

HasNameLabel returns a boolean if a field has been set.

### GetOptout

`func (o *WlanPortalTemplate) GetOptout() bool`

GetOptout returns the Optout field if non-nil, zero value otherwise.

### GetOptoutOk

`func (o *WlanPortalTemplate) GetOptoutOk() (*bool, bool)`

GetOptoutOk returns a tuple with the Optout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptout

`func (o *WlanPortalTemplate) SetOptout(v bool)`

SetOptout sets Optout field to given value.

### HasOptout

`func (o *WlanPortalTemplate) HasOptout() bool`

HasOptout returns a boolean if a field has been set.

### GetOptoutDefault

`func (o *WlanPortalTemplate) GetOptoutDefault() bool`

GetOptoutDefault returns the OptoutDefault field if non-nil, zero value otherwise.

### GetOptoutDefaultOk

`func (o *WlanPortalTemplate) GetOptoutDefaultOk() (*bool, bool)`

GetOptoutDefaultOk returns a tuple with the OptoutDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptoutDefault

`func (o *WlanPortalTemplate) SetOptoutDefault(v bool)`

SetOptoutDefault sets OptoutDefault field to given value.

### HasOptoutDefault

`func (o *WlanPortalTemplate) HasOptoutDefault() bool`

HasOptoutDefault returns a boolean if a field has been set.

### GetOptoutLabel

`func (o *WlanPortalTemplate) GetOptoutLabel() string`

GetOptoutLabel returns the OptoutLabel field if non-nil, zero value otherwise.

### GetOptoutLabelOk

`func (o *WlanPortalTemplate) GetOptoutLabelOk() (*string, bool)`

GetOptoutLabelOk returns a tuple with the OptoutLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptoutLabel

`func (o *WlanPortalTemplate) SetOptoutLabel(v string)`

SetOptoutLabel sets OptoutLabel field to given value.

### HasOptoutLabel

`func (o *WlanPortalTemplate) HasOptoutLabel() bool`

HasOptoutLabel returns a boolean if a field has been set.

### GetPageTitle

`func (o *WlanPortalTemplate) GetPageTitle() string`

GetPageTitle returns the PageTitle field if non-nil, zero value otherwise.

### GetPageTitleOk

`func (o *WlanPortalTemplate) GetPageTitleOk() (*string, bool)`

GetPageTitleOk returns a tuple with the PageTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageTitle

`func (o *WlanPortalTemplate) SetPageTitle(v string)`

SetPageTitle sets PageTitle field to given value.


### GetPassphraseCancel

`func (o *WlanPortalTemplate) GetPassphraseCancel() string`

GetPassphraseCancel returns the PassphraseCancel field if non-nil, zero value otherwise.

### GetPassphraseCancelOk

`func (o *WlanPortalTemplate) GetPassphraseCancelOk() (*string, bool)`

GetPassphraseCancelOk returns a tuple with the PassphraseCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseCancel

`func (o *WlanPortalTemplate) SetPassphraseCancel(v string)`

SetPassphraseCancel sets PassphraseCancel field to given value.

### HasPassphraseCancel

`func (o *WlanPortalTemplate) HasPassphraseCancel() bool`

HasPassphraseCancel returns a boolean if a field has been set.

### GetPassphraseError

`func (o *WlanPortalTemplate) GetPassphraseError() string`

GetPassphraseError returns the PassphraseError field if non-nil, zero value otherwise.

### GetPassphraseErrorOk

`func (o *WlanPortalTemplate) GetPassphraseErrorOk() (*string, bool)`

GetPassphraseErrorOk returns a tuple with the PassphraseError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseError

`func (o *WlanPortalTemplate) SetPassphraseError(v string)`

SetPassphraseError sets PassphraseError field to given value.

### HasPassphraseError

`func (o *WlanPortalTemplate) HasPassphraseError() bool`

HasPassphraseError returns a boolean if a field has been set.

### GetPassphraseLabel

`func (o *WlanPortalTemplate) GetPassphraseLabel() string`

GetPassphraseLabel returns the PassphraseLabel field if non-nil, zero value otherwise.

### GetPassphraseLabelOk

`func (o *WlanPortalTemplate) GetPassphraseLabelOk() (*string, bool)`

GetPassphraseLabelOk returns a tuple with the PassphraseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseLabel

`func (o *WlanPortalTemplate) SetPassphraseLabel(v string)`

SetPassphraseLabel sets PassphraseLabel field to given value.

### HasPassphraseLabel

`func (o *WlanPortalTemplate) HasPassphraseLabel() bool`

HasPassphraseLabel returns a boolean if a field has been set.

### GetPassphraseMessage

`func (o *WlanPortalTemplate) GetPassphraseMessage() string`

GetPassphraseMessage returns the PassphraseMessage field if non-nil, zero value otherwise.

### GetPassphraseMessageOk

`func (o *WlanPortalTemplate) GetPassphraseMessageOk() (*string, bool)`

GetPassphraseMessageOk returns a tuple with the PassphraseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseMessage

`func (o *WlanPortalTemplate) SetPassphraseMessage(v string)`

SetPassphraseMessage sets PassphraseMessage field to given value.

### HasPassphraseMessage

`func (o *WlanPortalTemplate) HasPassphraseMessage() bool`

HasPassphraseMessage returns a boolean if a field has been set.

### GetPassphraseSubmit

`func (o *WlanPortalTemplate) GetPassphraseSubmit() string`

GetPassphraseSubmit returns the PassphraseSubmit field if non-nil, zero value otherwise.

### GetPassphraseSubmitOk

`func (o *WlanPortalTemplate) GetPassphraseSubmitOk() (*string, bool)`

GetPassphraseSubmitOk returns a tuple with the PassphraseSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseSubmit

`func (o *WlanPortalTemplate) SetPassphraseSubmit(v string)`

SetPassphraseSubmit sets PassphraseSubmit field to given value.

### HasPassphraseSubmit

`func (o *WlanPortalTemplate) HasPassphraseSubmit() bool`

HasPassphraseSubmit returns a boolean if a field has been set.

### GetPassphraseTitle

`func (o *WlanPortalTemplate) GetPassphraseTitle() string`

GetPassphraseTitle returns the PassphraseTitle field if non-nil, zero value otherwise.

### GetPassphraseTitleOk

`func (o *WlanPortalTemplate) GetPassphraseTitleOk() (*string, bool)`

GetPassphraseTitleOk returns a tuple with the PassphraseTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseTitle

`func (o *WlanPortalTemplate) SetPassphraseTitle(v string)`

SetPassphraseTitle sets PassphraseTitle field to given value.

### HasPassphraseTitle

`func (o *WlanPortalTemplate) HasPassphraseTitle() bool`

HasPassphraseTitle returns a boolean if a field has been set.

### GetPoweredBy

`func (o *WlanPortalTemplate) GetPoweredBy() bool`

GetPoweredBy returns the PoweredBy field if non-nil, zero value otherwise.

### GetPoweredByOk

`func (o *WlanPortalTemplate) GetPoweredByOk() (*bool, bool)`

GetPoweredByOk returns a tuple with the PoweredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoweredBy

`func (o *WlanPortalTemplate) SetPoweredBy(v bool)`

SetPoweredBy sets PoweredBy field to given value.

### HasPoweredBy

`func (o *WlanPortalTemplate) HasPoweredBy() bool`

HasPoweredBy returns a boolean if a field has been set.

### GetRequiredFieldLabel

`func (o *WlanPortalTemplate) GetRequiredFieldLabel() string`

GetRequiredFieldLabel returns the RequiredFieldLabel field if non-nil, zero value otherwise.

### GetRequiredFieldLabelOk

`func (o *WlanPortalTemplate) GetRequiredFieldLabelOk() (*string, bool)`

GetRequiredFieldLabelOk returns a tuple with the RequiredFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFieldLabel

`func (o *WlanPortalTemplate) SetRequiredFieldLabel(v string)`

SetRequiredFieldLabel sets RequiredFieldLabel field to given value.

### HasRequiredFieldLabel

`func (o *WlanPortalTemplate) HasRequiredFieldLabel() bool`

HasRequiredFieldLabel returns a boolean if a field has been set.

### GetSignInLabel

`func (o *WlanPortalTemplate) GetSignInLabel() string`

GetSignInLabel returns the SignInLabel field if non-nil, zero value otherwise.

### GetSignInLabelOk

`func (o *WlanPortalTemplate) GetSignInLabelOk() (*string, bool)`

GetSignInLabelOk returns a tuple with the SignInLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInLabel

`func (o *WlanPortalTemplate) SetSignInLabel(v string)`

SetSignInLabel sets SignInLabel field to given value.

### HasSignInLabel

`func (o *WlanPortalTemplate) HasSignInLabel() bool`

HasSignInLabel returns a boolean if a field has been set.

### GetSmsCarrierDefault

`func (o *WlanPortalTemplate) GetSmsCarrierDefault() string`

GetSmsCarrierDefault returns the SmsCarrierDefault field if non-nil, zero value otherwise.

### GetSmsCarrierDefaultOk

`func (o *WlanPortalTemplate) GetSmsCarrierDefaultOk() (*string, bool)`

GetSmsCarrierDefaultOk returns a tuple with the SmsCarrierDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCarrierDefault

`func (o *WlanPortalTemplate) SetSmsCarrierDefault(v string)`

SetSmsCarrierDefault sets SmsCarrierDefault field to given value.

### HasSmsCarrierDefault

`func (o *WlanPortalTemplate) HasSmsCarrierDefault() bool`

HasSmsCarrierDefault returns a boolean if a field has been set.

### GetSmsCarrierError

`func (o *WlanPortalTemplate) GetSmsCarrierError() string`

GetSmsCarrierError returns the SmsCarrierError field if non-nil, zero value otherwise.

### GetSmsCarrierErrorOk

`func (o *WlanPortalTemplate) GetSmsCarrierErrorOk() (*string, bool)`

GetSmsCarrierErrorOk returns a tuple with the SmsCarrierError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCarrierError

`func (o *WlanPortalTemplate) SetSmsCarrierError(v string)`

SetSmsCarrierError sets SmsCarrierError field to given value.

### HasSmsCarrierError

`func (o *WlanPortalTemplate) HasSmsCarrierError() bool`

HasSmsCarrierError returns a boolean if a field has been set.

### GetSmsCarrierFieldLabel

`func (o *WlanPortalTemplate) GetSmsCarrierFieldLabel() string`

GetSmsCarrierFieldLabel returns the SmsCarrierFieldLabel field if non-nil, zero value otherwise.

### GetSmsCarrierFieldLabelOk

`func (o *WlanPortalTemplate) GetSmsCarrierFieldLabelOk() (*string, bool)`

GetSmsCarrierFieldLabelOk returns a tuple with the SmsCarrierFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCarrierFieldLabel

`func (o *WlanPortalTemplate) SetSmsCarrierFieldLabel(v string)`

SetSmsCarrierFieldLabel sets SmsCarrierFieldLabel field to given value.

### HasSmsCarrierFieldLabel

`func (o *WlanPortalTemplate) HasSmsCarrierFieldLabel() bool`

HasSmsCarrierFieldLabel returns a boolean if a field has been set.

### GetSmsCodeCancel

`func (o *WlanPortalTemplate) GetSmsCodeCancel() string`

GetSmsCodeCancel returns the SmsCodeCancel field if non-nil, zero value otherwise.

### GetSmsCodeCancelOk

`func (o *WlanPortalTemplate) GetSmsCodeCancelOk() (*string, bool)`

GetSmsCodeCancelOk returns a tuple with the SmsCodeCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCodeCancel

`func (o *WlanPortalTemplate) SetSmsCodeCancel(v string)`

SetSmsCodeCancel sets SmsCodeCancel field to given value.

### HasSmsCodeCancel

`func (o *WlanPortalTemplate) HasSmsCodeCancel() bool`

HasSmsCodeCancel returns a boolean if a field has been set.

### GetSmsCodeError

`func (o *WlanPortalTemplate) GetSmsCodeError() string`

GetSmsCodeError returns the SmsCodeError field if non-nil, zero value otherwise.

### GetSmsCodeErrorOk

`func (o *WlanPortalTemplate) GetSmsCodeErrorOk() (*string, bool)`

GetSmsCodeErrorOk returns a tuple with the SmsCodeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCodeError

`func (o *WlanPortalTemplate) SetSmsCodeError(v string)`

SetSmsCodeError sets SmsCodeError field to given value.

### HasSmsCodeError

`func (o *WlanPortalTemplate) HasSmsCodeError() bool`

HasSmsCodeError returns a boolean if a field has been set.

### GetSmsCodeFieldLabel

`func (o *WlanPortalTemplate) GetSmsCodeFieldLabel() string`

GetSmsCodeFieldLabel returns the SmsCodeFieldLabel field if non-nil, zero value otherwise.

### GetSmsCodeFieldLabelOk

`func (o *WlanPortalTemplate) GetSmsCodeFieldLabelOk() (*string, bool)`

GetSmsCodeFieldLabelOk returns a tuple with the SmsCodeFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCodeFieldLabel

`func (o *WlanPortalTemplate) SetSmsCodeFieldLabel(v string)`

SetSmsCodeFieldLabel sets SmsCodeFieldLabel field to given value.

### HasSmsCodeFieldLabel

`func (o *WlanPortalTemplate) HasSmsCodeFieldLabel() bool`

HasSmsCodeFieldLabel returns a boolean if a field has been set.

### GetSmsCodeMessage

`func (o *WlanPortalTemplate) GetSmsCodeMessage() string`

GetSmsCodeMessage returns the SmsCodeMessage field if non-nil, zero value otherwise.

### GetSmsCodeMessageOk

`func (o *WlanPortalTemplate) GetSmsCodeMessageOk() (*string, bool)`

GetSmsCodeMessageOk returns a tuple with the SmsCodeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCodeMessage

`func (o *WlanPortalTemplate) SetSmsCodeMessage(v string)`

SetSmsCodeMessage sets SmsCodeMessage field to given value.

### HasSmsCodeMessage

`func (o *WlanPortalTemplate) HasSmsCodeMessage() bool`

HasSmsCodeMessage returns a boolean if a field has been set.

### GetSmsCodeSubmit

`func (o *WlanPortalTemplate) GetSmsCodeSubmit() string`

GetSmsCodeSubmit returns the SmsCodeSubmit field if non-nil, zero value otherwise.

### GetSmsCodeSubmitOk

`func (o *WlanPortalTemplate) GetSmsCodeSubmitOk() (*string, bool)`

GetSmsCodeSubmitOk returns a tuple with the SmsCodeSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCodeSubmit

`func (o *WlanPortalTemplate) SetSmsCodeSubmit(v string)`

SetSmsCodeSubmit sets SmsCodeSubmit field to given value.

### HasSmsCodeSubmit

`func (o *WlanPortalTemplate) HasSmsCodeSubmit() bool`

HasSmsCodeSubmit returns a boolean if a field has been set.

### GetSmsCodeTitle

`func (o *WlanPortalTemplate) GetSmsCodeTitle() string`

GetSmsCodeTitle returns the SmsCodeTitle field if non-nil, zero value otherwise.

### GetSmsCodeTitleOk

`func (o *WlanPortalTemplate) GetSmsCodeTitleOk() (*string, bool)`

GetSmsCodeTitleOk returns a tuple with the SmsCodeTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCodeTitle

`func (o *WlanPortalTemplate) SetSmsCodeTitle(v string)`

SetSmsCodeTitle sets SmsCodeTitle field to given value.

### HasSmsCodeTitle

`func (o *WlanPortalTemplate) HasSmsCodeTitle() bool`

HasSmsCodeTitle returns a boolean if a field has been set.

### GetSmsCountryFieldLabel

`func (o *WlanPortalTemplate) GetSmsCountryFieldLabel() string`

GetSmsCountryFieldLabel returns the SmsCountryFieldLabel field if non-nil, zero value otherwise.

### GetSmsCountryFieldLabelOk

`func (o *WlanPortalTemplate) GetSmsCountryFieldLabelOk() (*string, bool)`

GetSmsCountryFieldLabelOk returns a tuple with the SmsCountryFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCountryFieldLabel

`func (o *WlanPortalTemplate) SetSmsCountryFieldLabel(v string)`

SetSmsCountryFieldLabel sets SmsCountryFieldLabel field to given value.

### HasSmsCountryFieldLabel

`func (o *WlanPortalTemplate) HasSmsCountryFieldLabel() bool`

HasSmsCountryFieldLabel returns a boolean if a field has been set.

### GetSmsCountryFormat

`func (o *WlanPortalTemplate) GetSmsCountryFormat() string`

GetSmsCountryFormat returns the SmsCountryFormat field if non-nil, zero value otherwise.

### GetSmsCountryFormatOk

`func (o *WlanPortalTemplate) GetSmsCountryFormatOk() (*string, bool)`

GetSmsCountryFormatOk returns a tuple with the SmsCountryFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCountryFormat

`func (o *WlanPortalTemplate) SetSmsCountryFormat(v string)`

SetSmsCountryFormat sets SmsCountryFormat field to given value.

### HasSmsCountryFormat

`func (o *WlanPortalTemplate) HasSmsCountryFormat() bool`

HasSmsCountryFormat returns a boolean if a field has been set.

### GetSmsHaveAccessCode

`func (o *WlanPortalTemplate) GetSmsHaveAccessCode() string`

GetSmsHaveAccessCode returns the SmsHaveAccessCode field if non-nil, zero value otherwise.

### GetSmsHaveAccessCodeOk

`func (o *WlanPortalTemplate) GetSmsHaveAccessCodeOk() (*string, bool)`

GetSmsHaveAccessCodeOk returns a tuple with the SmsHaveAccessCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsHaveAccessCode

`func (o *WlanPortalTemplate) SetSmsHaveAccessCode(v string)`

SetSmsHaveAccessCode sets SmsHaveAccessCode field to given value.

### HasSmsHaveAccessCode

`func (o *WlanPortalTemplate) HasSmsHaveAccessCode() bool`

HasSmsHaveAccessCode returns a boolean if a field has been set.

### GetSmsMessageFormat

`func (o *WlanPortalTemplate) GetSmsMessageFormat() string`

GetSmsMessageFormat returns the SmsMessageFormat field if non-nil, zero value otherwise.

### GetSmsMessageFormatOk

`func (o *WlanPortalTemplate) GetSmsMessageFormatOk() (*string, bool)`

GetSmsMessageFormatOk returns a tuple with the SmsMessageFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMessageFormat

`func (o *WlanPortalTemplate) SetSmsMessageFormat(v string)`

SetSmsMessageFormat sets SmsMessageFormat field to given value.

### HasSmsMessageFormat

`func (o *WlanPortalTemplate) HasSmsMessageFormat() bool`

HasSmsMessageFormat returns a boolean if a field has been set.

### GetSmsNumberCancel

`func (o *WlanPortalTemplate) GetSmsNumberCancel() string`

GetSmsNumberCancel returns the SmsNumberCancel field if non-nil, zero value otherwise.

### GetSmsNumberCancelOk

`func (o *WlanPortalTemplate) GetSmsNumberCancelOk() (*string, bool)`

GetSmsNumberCancelOk returns a tuple with the SmsNumberCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumberCancel

`func (o *WlanPortalTemplate) SetSmsNumberCancel(v string)`

SetSmsNumberCancel sets SmsNumberCancel field to given value.

### HasSmsNumberCancel

`func (o *WlanPortalTemplate) HasSmsNumberCancel() bool`

HasSmsNumberCancel returns a boolean if a field has been set.

### GetSmsNumberError

`func (o *WlanPortalTemplate) GetSmsNumberError() string`

GetSmsNumberError returns the SmsNumberError field if non-nil, zero value otherwise.

### GetSmsNumberErrorOk

`func (o *WlanPortalTemplate) GetSmsNumberErrorOk() (*string, bool)`

GetSmsNumberErrorOk returns a tuple with the SmsNumberError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumberError

`func (o *WlanPortalTemplate) SetSmsNumberError(v string)`

SetSmsNumberError sets SmsNumberError field to given value.

### HasSmsNumberError

`func (o *WlanPortalTemplate) HasSmsNumberError() bool`

HasSmsNumberError returns a boolean if a field has been set.

### GetSmsNumberFieldLabel

`func (o *WlanPortalTemplate) GetSmsNumberFieldLabel() string`

GetSmsNumberFieldLabel returns the SmsNumberFieldLabel field if non-nil, zero value otherwise.

### GetSmsNumberFieldLabelOk

`func (o *WlanPortalTemplate) GetSmsNumberFieldLabelOk() (*string, bool)`

GetSmsNumberFieldLabelOk returns a tuple with the SmsNumberFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumberFieldLabel

`func (o *WlanPortalTemplate) SetSmsNumberFieldLabel(v string)`

SetSmsNumberFieldLabel sets SmsNumberFieldLabel field to given value.

### HasSmsNumberFieldLabel

`func (o *WlanPortalTemplate) HasSmsNumberFieldLabel() bool`

HasSmsNumberFieldLabel returns a boolean if a field has been set.

### GetSmsNumberFormat

`func (o *WlanPortalTemplate) GetSmsNumberFormat() string`

GetSmsNumberFormat returns the SmsNumberFormat field if non-nil, zero value otherwise.

### GetSmsNumberFormatOk

`func (o *WlanPortalTemplate) GetSmsNumberFormatOk() (*string, bool)`

GetSmsNumberFormatOk returns a tuple with the SmsNumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumberFormat

`func (o *WlanPortalTemplate) SetSmsNumberFormat(v string)`

SetSmsNumberFormat sets SmsNumberFormat field to given value.

### HasSmsNumberFormat

`func (o *WlanPortalTemplate) HasSmsNumberFormat() bool`

HasSmsNumberFormat returns a boolean if a field has been set.

### GetSmsNumberMessage

`func (o *WlanPortalTemplate) GetSmsNumberMessage() string`

GetSmsNumberMessage returns the SmsNumberMessage field if non-nil, zero value otherwise.

### GetSmsNumberMessageOk

`func (o *WlanPortalTemplate) GetSmsNumberMessageOk() (*string, bool)`

GetSmsNumberMessageOk returns a tuple with the SmsNumberMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumberMessage

`func (o *WlanPortalTemplate) SetSmsNumberMessage(v string)`

SetSmsNumberMessage sets SmsNumberMessage field to given value.

### HasSmsNumberMessage

`func (o *WlanPortalTemplate) HasSmsNumberMessage() bool`

HasSmsNumberMessage returns a boolean if a field has been set.

### GetSmsNumberSubmit

`func (o *WlanPortalTemplate) GetSmsNumberSubmit() string`

GetSmsNumberSubmit returns the SmsNumberSubmit field if non-nil, zero value otherwise.

### GetSmsNumberSubmitOk

`func (o *WlanPortalTemplate) GetSmsNumberSubmitOk() (*string, bool)`

GetSmsNumberSubmitOk returns a tuple with the SmsNumberSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumberSubmit

`func (o *WlanPortalTemplate) SetSmsNumberSubmit(v string)`

SetSmsNumberSubmit sets SmsNumberSubmit field to given value.

### HasSmsNumberSubmit

`func (o *WlanPortalTemplate) HasSmsNumberSubmit() bool`

HasSmsNumberSubmit returns a boolean if a field has been set.

### GetSmsNumberTitle

`func (o *WlanPortalTemplate) GetSmsNumberTitle() string`

GetSmsNumberTitle returns the SmsNumberTitle field if non-nil, zero value otherwise.

### GetSmsNumberTitleOk

`func (o *WlanPortalTemplate) GetSmsNumberTitleOk() (*string, bool)`

GetSmsNumberTitleOk returns a tuple with the SmsNumberTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumberTitle

`func (o *WlanPortalTemplate) SetSmsNumberTitle(v string)`

SetSmsNumberTitle sets SmsNumberTitle field to given value.

### HasSmsNumberTitle

`func (o *WlanPortalTemplate) HasSmsNumberTitle() bool`

HasSmsNumberTitle returns a boolean if a field has been set.

### GetSmsUsernameFormat

`func (o *WlanPortalTemplate) GetSmsUsernameFormat() string`

GetSmsUsernameFormat returns the SmsUsernameFormat field if non-nil, zero value otherwise.

### GetSmsUsernameFormatOk

`func (o *WlanPortalTemplate) GetSmsUsernameFormatOk() (*string, bool)`

GetSmsUsernameFormatOk returns a tuple with the SmsUsernameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsUsernameFormat

`func (o *WlanPortalTemplate) SetSmsUsernameFormat(v string)`

SetSmsUsernameFormat sets SmsUsernameFormat field to given value.

### HasSmsUsernameFormat

`func (o *WlanPortalTemplate) HasSmsUsernameFormat() bool`

HasSmsUsernameFormat returns a boolean if a field has been set.

### GetSmsValidityDuration

`func (o *WlanPortalTemplate) GetSmsValidityDuration() int32`

GetSmsValidityDuration returns the SmsValidityDuration field if non-nil, zero value otherwise.

### GetSmsValidityDurationOk

`func (o *WlanPortalTemplate) GetSmsValidityDurationOk() (*int32, bool)`

GetSmsValidityDurationOk returns a tuple with the SmsValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsValidityDuration

`func (o *WlanPortalTemplate) SetSmsValidityDuration(v int32)`

SetSmsValidityDuration sets SmsValidityDuration field to given value.

### HasSmsValidityDuration

`func (o *WlanPortalTemplate) HasSmsValidityDuration() bool`

HasSmsValidityDuration returns a boolean if a field has been set.

### GetSponsorBackLink

`func (o *WlanPortalTemplate) GetSponsorBackLink() string`

GetSponsorBackLink returns the SponsorBackLink field if non-nil, zero value otherwise.

### GetSponsorBackLinkOk

`func (o *WlanPortalTemplate) GetSponsorBackLinkOk() (*string, bool)`

GetSponsorBackLinkOk returns a tuple with the SponsorBackLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorBackLink

`func (o *WlanPortalTemplate) SetSponsorBackLink(v string)`

SetSponsorBackLink sets SponsorBackLink field to given value.

### HasSponsorBackLink

`func (o *WlanPortalTemplate) HasSponsorBackLink() bool`

HasSponsorBackLink returns a boolean if a field has been set.

### GetSponsorCancel

`func (o *WlanPortalTemplate) GetSponsorCancel() string`

GetSponsorCancel returns the SponsorCancel field if non-nil, zero value otherwise.

### GetSponsorCancelOk

`func (o *WlanPortalTemplate) GetSponsorCancelOk() (*string, bool)`

GetSponsorCancelOk returns a tuple with the SponsorCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorCancel

`func (o *WlanPortalTemplate) SetSponsorCancel(v string)`

SetSponsorCancel sets SponsorCancel field to given value.

### HasSponsorCancel

`func (o *WlanPortalTemplate) HasSponsorCancel() bool`

HasSponsorCancel returns a boolean if a field has been set.

### GetSponsorEmail

`func (o *WlanPortalTemplate) GetSponsorEmail() string`

GetSponsorEmail returns the SponsorEmail field if non-nil, zero value otherwise.

### GetSponsorEmailOk

`func (o *WlanPortalTemplate) GetSponsorEmailOk() (*string, bool)`

GetSponsorEmailOk returns a tuple with the SponsorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorEmail

`func (o *WlanPortalTemplate) SetSponsorEmail(v string)`

SetSponsorEmail sets SponsorEmail field to given value.

### HasSponsorEmail

`func (o *WlanPortalTemplate) HasSponsorEmail() bool`

HasSponsorEmail returns a boolean if a field has been set.

### GetSponsorEmailError

`func (o *WlanPortalTemplate) GetSponsorEmailError() string`

GetSponsorEmailError returns the SponsorEmailError field if non-nil, zero value otherwise.

### GetSponsorEmailErrorOk

`func (o *WlanPortalTemplate) GetSponsorEmailErrorOk() (*string, bool)`

GetSponsorEmailErrorOk returns a tuple with the SponsorEmailError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorEmailError

`func (o *WlanPortalTemplate) SetSponsorEmailError(v string)`

SetSponsorEmailError sets SponsorEmailError field to given value.

### HasSponsorEmailError

`func (o *WlanPortalTemplate) HasSponsorEmailError() bool`

HasSponsorEmailError returns a boolean if a field has been set.

### GetSponsorEmailTemplate

`func (o *WlanPortalTemplate) GetSponsorEmailTemplate() string`

GetSponsorEmailTemplate returns the SponsorEmailTemplate field if non-nil, zero value otherwise.

### GetSponsorEmailTemplateOk

`func (o *WlanPortalTemplate) GetSponsorEmailTemplateOk() (*string, bool)`

GetSponsorEmailTemplateOk returns a tuple with the SponsorEmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorEmailTemplate

`func (o *WlanPortalTemplate) SetSponsorEmailTemplate(v string)`

SetSponsorEmailTemplate sets SponsorEmailTemplate field to given value.

### HasSponsorEmailTemplate

`func (o *WlanPortalTemplate) HasSponsorEmailTemplate() bool`

HasSponsorEmailTemplate returns a boolean if a field has been set.

### GetSponsorInfoApproved

`func (o *WlanPortalTemplate) GetSponsorInfoApproved() string`

GetSponsorInfoApproved returns the SponsorInfoApproved field if non-nil, zero value otherwise.

### GetSponsorInfoApprovedOk

`func (o *WlanPortalTemplate) GetSponsorInfoApprovedOk() (*string, bool)`

GetSponsorInfoApprovedOk returns a tuple with the SponsorInfoApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorInfoApproved

`func (o *WlanPortalTemplate) SetSponsorInfoApproved(v string)`

SetSponsorInfoApproved sets SponsorInfoApproved field to given value.

### HasSponsorInfoApproved

`func (o *WlanPortalTemplate) HasSponsorInfoApproved() bool`

HasSponsorInfoApproved returns a boolean if a field has been set.

### GetSponsorInfoDenied

`func (o *WlanPortalTemplate) GetSponsorInfoDenied() string`

GetSponsorInfoDenied returns the SponsorInfoDenied field if non-nil, zero value otherwise.

### GetSponsorInfoDeniedOk

`func (o *WlanPortalTemplate) GetSponsorInfoDeniedOk() (*string, bool)`

GetSponsorInfoDeniedOk returns a tuple with the SponsorInfoDenied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorInfoDenied

`func (o *WlanPortalTemplate) SetSponsorInfoDenied(v string)`

SetSponsorInfoDenied sets SponsorInfoDenied field to given value.

### HasSponsorInfoDenied

`func (o *WlanPortalTemplate) HasSponsorInfoDenied() bool`

HasSponsorInfoDenied returns a boolean if a field has been set.

### GetSponsorInfoPending

`func (o *WlanPortalTemplate) GetSponsorInfoPending() string`

GetSponsorInfoPending returns the SponsorInfoPending field if non-nil, zero value otherwise.

### GetSponsorInfoPendingOk

`func (o *WlanPortalTemplate) GetSponsorInfoPendingOk() (*string, bool)`

GetSponsorInfoPendingOk returns a tuple with the SponsorInfoPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorInfoPending

`func (o *WlanPortalTemplate) SetSponsorInfoPending(v string)`

SetSponsorInfoPending sets SponsorInfoPending field to given value.

### HasSponsorInfoPending

`func (o *WlanPortalTemplate) HasSponsorInfoPending() bool`

HasSponsorInfoPending returns a boolean if a field has been set.

### GetSponsorName

`func (o *WlanPortalTemplate) GetSponsorName() string`

GetSponsorName returns the SponsorName field if non-nil, zero value otherwise.

### GetSponsorNameOk

`func (o *WlanPortalTemplate) GetSponsorNameOk() (*string, bool)`

GetSponsorNameOk returns a tuple with the SponsorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorName

`func (o *WlanPortalTemplate) SetSponsorName(v string)`

SetSponsorName sets SponsorName field to given value.

### HasSponsorName

`func (o *WlanPortalTemplate) HasSponsorName() bool`

HasSponsorName returns a boolean if a field has been set.

### GetSponsorNameError

`func (o *WlanPortalTemplate) GetSponsorNameError() string`

GetSponsorNameError returns the SponsorNameError field if non-nil, zero value otherwise.

### GetSponsorNameErrorOk

`func (o *WlanPortalTemplate) GetSponsorNameErrorOk() (*string, bool)`

GetSponsorNameErrorOk returns a tuple with the SponsorNameError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorNameError

`func (o *WlanPortalTemplate) SetSponsorNameError(v string)`

SetSponsorNameError sets SponsorNameError field to given value.

### HasSponsorNameError

`func (o *WlanPortalTemplate) HasSponsorNameError() bool`

HasSponsorNameError returns a boolean if a field has been set.

### GetSponsorNotePending

`func (o *WlanPortalTemplate) GetSponsorNotePending() string`

GetSponsorNotePending returns the SponsorNotePending field if non-nil, zero value otherwise.

### GetSponsorNotePendingOk

`func (o *WlanPortalTemplate) GetSponsorNotePendingOk() (*string, bool)`

GetSponsorNotePendingOk returns a tuple with the SponsorNotePending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorNotePending

`func (o *WlanPortalTemplate) SetSponsorNotePending(v string)`

SetSponsorNotePending sets SponsorNotePending field to given value.

### HasSponsorNotePending

`func (o *WlanPortalTemplate) HasSponsorNotePending() bool`

HasSponsorNotePending returns a boolean if a field has been set.

### GetSponsorRequestAccess

`func (o *WlanPortalTemplate) GetSponsorRequestAccess() string`

GetSponsorRequestAccess returns the SponsorRequestAccess field if non-nil, zero value otherwise.

### GetSponsorRequestAccessOk

`func (o *WlanPortalTemplate) GetSponsorRequestAccessOk() (*string, bool)`

GetSponsorRequestAccessOk returns a tuple with the SponsorRequestAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorRequestAccess

`func (o *WlanPortalTemplate) SetSponsorRequestAccess(v string)`

SetSponsorRequestAccess sets SponsorRequestAccess field to given value.

### HasSponsorRequestAccess

`func (o *WlanPortalTemplate) HasSponsorRequestAccess() bool`

HasSponsorRequestAccess returns a boolean if a field has been set.

### GetSponsorSelectEmail

`func (o *WlanPortalTemplate) GetSponsorSelectEmail() string`

GetSponsorSelectEmail returns the SponsorSelectEmail field if non-nil, zero value otherwise.

### GetSponsorSelectEmailOk

`func (o *WlanPortalTemplate) GetSponsorSelectEmailOk() (*string, bool)`

GetSponsorSelectEmailOk returns a tuple with the SponsorSelectEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorSelectEmail

`func (o *WlanPortalTemplate) SetSponsorSelectEmail(v string)`

SetSponsorSelectEmail sets SponsorSelectEmail field to given value.

### HasSponsorSelectEmail

`func (o *WlanPortalTemplate) HasSponsorSelectEmail() bool`

HasSponsorSelectEmail returns a boolean if a field has been set.

### GetSponsorStatusApproved

`func (o *WlanPortalTemplate) GetSponsorStatusApproved() string`

GetSponsorStatusApproved returns the SponsorStatusApproved field if non-nil, zero value otherwise.

### GetSponsorStatusApprovedOk

`func (o *WlanPortalTemplate) GetSponsorStatusApprovedOk() (*string, bool)`

GetSponsorStatusApprovedOk returns a tuple with the SponsorStatusApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorStatusApproved

`func (o *WlanPortalTemplate) SetSponsorStatusApproved(v string)`

SetSponsorStatusApproved sets SponsorStatusApproved field to given value.

### HasSponsorStatusApproved

`func (o *WlanPortalTemplate) HasSponsorStatusApproved() bool`

HasSponsorStatusApproved returns a boolean if a field has been set.

### GetSponsorStatusDenied

`func (o *WlanPortalTemplate) GetSponsorStatusDenied() string`

GetSponsorStatusDenied returns the SponsorStatusDenied field if non-nil, zero value otherwise.

### GetSponsorStatusDeniedOk

`func (o *WlanPortalTemplate) GetSponsorStatusDeniedOk() (*string, bool)`

GetSponsorStatusDeniedOk returns a tuple with the SponsorStatusDenied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorStatusDenied

`func (o *WlanPortalTemplate) SetSponsorStatusDenied(v string)`

SetSponsorStatusDenied sets SponsorStatusDenied field to given value.

### HasSponsorStatusDenied

`func (o *WlanPortalTemplate) HasSponsorStatusDenied() bool`

HasSponsorStatusDenied returns a boolean if a field has been set.

### GetSponsorStatusPending

`func (o *WlanPortalTemplate) GetSponsorStatusPending() string`

GetSponsorStatusPending returns the SponsorStatusPending field if non-nil, zero value otherwise.

### GetSponsorStatusPendingOk

`func (o *WlanPortalTemplate) GetSponsorStatusPendingOk() (*string, bool)`

GetSponsorStatusPendingOk returns a tuple with the SponsorStatusPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorStatusPending

`func (o *WlanPortalTemplate) SetSponsorStatusPending(v string)`

SetSponsorStatusPending sets SponsorStatusPending field to given value.

### HasSponsorStatusPending

`func (o *WlanPortalTemplate) HasSponsorStatusPending() bool`

HasSponsorStatusPending returns a boolean if a field has been set.

### GetSponsorSubmit

`func (o *WlanPortalTemplate) GetSponsorSubmit() string`

GetSponsorSubmit returns the SponsorSubmit field if non-nil, zero value otherwise.

### GetSponsorSubmitOk

`func (o *WlanPortalTemplate) GetSponsorSubmitOk() (*string, bool)`

GetSponsorSubmitOk returns a tuple with the SponsorSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorSubmit

`func (o *WlanPortalTemplate) SetSponsorSubmit(v string)`

SetSponsorSubmit sets SponsorSubmit field to given value.

### HasSponsorSubmit

`func (o *WlanPortalTemplate) HasSponsorSubmit() bool`

HasSponsorSubmit returns a boolean if a field has been set.

### GetSponsorsError

`func (o *WlanPortalTemplate) GetSponsorsError() string`

GetSponsorsError returns the SponsorsError field if non-nil, zero value otherwise.

### GetSponsorsErrorOk

`func (o *WlanPortalTemplate) GetSponsorsErrorOk() (*string, bool)`

GetSponsorsErrorOk returns a tuple with the SponsorsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorsError

`func (o *WlanPortalTemplate) SetSponsorsError(v string)`

SetSponsorsError sets SponsorsError field to given value.

### HasSponsorsError

`func (o *WlanPortalTemplate) HasSponsorsError() bool`

HasSponsorsError returns a boolean if a field has been set.

### GetSponsorsInfoApproved

`func (o *WlanPortalTemplate) GetSponsorsInfoApproved() string`

GetSponsorsInfoApproved returns the SponsorsInfoApproved field if non-nil, zero value otherwise.

### GetSponsorsInfoApprovedOk

`func (o *WlanPortalTemplate) GetSponsorsInfoApprovedOk() (*string, bool)`

GetSponsorsInfoApprovedOk returns a tuple with the SponsorsInfoApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorsInfoApproved

`func (o *WlanPortalTemplate) SetSponsorsInfoApproved(v string)`

SetSponsorsInfoApproved sets SponsorsInfoApproved field to given value.

### HasSponsorsInfoApproved

`func (o *WlanPortalTemplate) HasSponsorsInfoApproved() bool`

HasSponsorsInfoApproved returns a boolean if a field has been set.

### GetSponsorsInfoDenied

`func (o *WlanPortalTemplate) GetSponsorsInfoDenied() string`

GetSponsorsInfoDenied returns the SponsorsInfoDenied field if non-nil, zero value otherwise.

### GetSponsorsInfoDeniedOk

`func (o *WlanPortalTemplate) GetSponsorsInfoDeniedOk() (*string, bool)`

GetSponsorsInfoDeniedOk returns a tuple with the SponsorsInfoDenied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorsInfoDenied

`func (o *WlanPortalTemplate) SetSponsorsInfoDenied(v string)`

SetSponsorsInfoDenied sets SponsorsInfoDenied field to given value.

### HasSponsorsInfoDenied

`func (o *WlanPortalTemplate) HasSponsorsInfoDenied() bool`

HasSponsorsInfoDenied returns a boolean if a field has been set.

### GetSponsorsInfoPending

`func (o *WlanPortalTemplate) GetSponsorsInfoPending() string`

GetSponsorsInfoPending returns the SponsorsInfoPending field if non-nil, zero value otherwise.

### GetSponsorsInfoPendingOk

`func (o *WlanPortalTemplate) GetSponsorsInfoPendingOk() (*string, bool)`

GetSponsorsInfoPendingOk returns a tuple with the SponsorsInfoPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorsInfoPending

`func (o *WlanPortalTemplate) SetSponsorsInfoPending(v string)`

SetSponsorsInfoPending sets SponsorsInfoPending field to given value.

### HasSponsorsInfoPending

`func (o *WlanPortalTemplate) HasSponsorsInfoPending() bool`

HasSponsorsInfoPending returns a boolean if a field has been set.

### GetTos

`func (o *WlanPortalTemplate) GetTos() bool`

GetTos returns the Tos field if non-nil, zero value otherwise.

### GetTosOk

`func (o *WlanPortalTemplate) GetTosOk() (*bool, bool)`

GetTosOk returns a tuple with the Tos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTos

`func (o *WlanPortalTemplate) SetTos(v bool)`

SetTos sets Tos field to given value.

### HasTos

`func (o *WlanPortalTemplate) HasTos() bool`

HasTos returns a boolean if a field has been set.

### GetTosAcceptLabel

`func (o *WlanPortalTemplate) GetTosAcceptLabel() string`

GetTosAcceptLabel returns the TosAcceptLabel field if non-nil, zero value otherwise.

### GetTosAcceptLabelOk

`func (o *WlanPortalTemplate) GetTosAcceptLabelOk() (*string, bool)`

GetTosAcceptLabelOk returns a tuple with the TosAcceptLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosAcceptLabel

`func (o *WlanPortalTemplate) SetTosAcceptLabel(v string)`

SetTosAcceptLabel sets TosAcceptLabel field to given value.

### HasTosAcceptLabel

`func (o *WlanPortalTemplate) HasTosAcceptLabel() bool`

HasTosAcceptLabel returns a boolean if a field has been set.

### GetTosError

`func (o *WlanPortalTemplate) GetTosError() string`

GetTosError returns the TosError field if non-nil, zero value otherwise.

### GetTosErrorOk

`func (o *WlanPortalTemplate) GetTosErrorOk() (*string, bool)`

GetTosErrorOk returns a tuple with the TosError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosError

`func (o *WlanPortalTemplate) SetTosError(v string)`

SetTosError sets TosError field to given value.

### HasTosError

`func (o *WlanPortalTemplate) HasTosError() bool`

HasTosError returns a boolean if a field has been set.

### GetTosLink

`func (o *WlanPortalTemplate) GetTosLink() string`

GetTosLink returns the TosLink field if non-nil, zero value otherwise.

### GetTosLinkOk

`func (o *WlanPortalTemplate) GetTosLinkOk() (*string, bool)`

GetTosLinkOk returns a tuple with the TosLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosLink

`func (o *WlanPortalTemplate) SetTosLink(v string)`

SetTosLink sets TosLink field to given value.

### HasTosLink

`func (o *WlanPortalTemplate) HasTosLink() bool`

HasTosLink returns a boolean if a field has been set.

### GetTosText

`func (o *WlanPortalTemplate) GetTosText() string`

GetTosText returns the TosText field if non-nil, zero value otherwise.

### GetTosTextOk

`func (o *WlanPortalTemplate) GetTosTextOk() (*string, bool)`

GetTosTextOk returns a tuple with the TosText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosText

`func (o *WlanPortalTemplate) SetTosText(v string)`

SetTosText sets TosText field to given value.

### HasTosText

`func (o *WlanPortalTemplate) HasTosText() bool`

HasTosText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


