<template>
    <svg width="557" height="556" viewBox="0 0 557 556" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path
            d="M556.229 278C556.229 431.535 431.764 556 278.229 556C124.693 556 0.228558 431.535 0.228558 278C0.228558 229.618 12.5878 184.123 34.3214 144.5C81.5574 58.3819 173.075 0 278.229 0C383.382 0 474.9 58.382 522.136 144.5C525.27 150.214 528.209 156.051 530.945 162C547.177 197.305 556.229 236.596 556.229 278Z"
            fill="#C0C0C0" fill-opacity="0.36" />
        <path
            d="M522.136 144.5C474.9 58.382 383.382 0 278.229 0C173.075 0 81.5574 58.3819 34.3214 144.5H278.229H522.136Z"
            fill="url(#paint0_linear_1_14)" />
        <path
        class="withStroke"
            d="M278.229 2.5C380.919 2.5 470.491 58.6835 517.876 142H278.229H38.5813C85.9661 58.6835 175.538 2.5 278.229 2.5ZM278.229 147H520.651C523.493 152.248 526.169 157.598 528.673 163.044C544.758 198.029 553.729 236.963 553.729 278C553.729 430.154 430.383 553.5 278.229 553.5C126.074 553.5 2.72856 430.154 2.72856 278C2.72856 230.577 14.7088 185.959 35.806 147H278.229Z"
            stroke="black" stroke-width="5" />
        <path
        class="withStroke"
            d="M332.229 296.297C335.729 287.797 322.229 236.797 306.229 234.797C292.229 225.797 179.162 295.864 171.729 303.297C159.729 315.297 164.729 318.297 179.229 318.297C188.448 318.297 408.729 324.797 406.229 318.297C403.729 311.797 348.729 303.297 332.229 296.297Z"
            fill="url(#paint1_linear_1_14)" stroke="black" stroke-width="5" />
        <defs>
            <linearGradient id="paint0_linear_1_14" x1="106.321" y1="21.5" x2="555.821" y2="154.5"
                gradientUnits="userSpaceOnUse">
                <stop class="stoptop1" stop-color="#939393" />
                <stop class="stoptop2" offset="1" stop-color="white" />
            </linearGradient>
            <linearGradient id="paint1_linear_1_14" x1="285.625" y1="234" x2="331.729" y2="348.297"
                gradientUnits="userSpaceOnUse">
                <stop class="stop1" stop-color="#999999" />
                <stop class="stop2" offset="1" stop-color="#DADADA" />
            </linearGradient>
        </defs>
    </svg>
</template>

<style lang="sass" scoped>
    svg
        transition: all .3s ease-in
        *
            transition: all .3s ease-in
        &:hover
            .withStroke
                stroke: #999
            .stop1
                stop-color: blue
            .stoptop1
                stop-color: red
            .stoptop2
                stop-color: blue
            .stop2
                stop-color: red

            
        
</style>