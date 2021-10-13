<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="6">
      <v-card>
        <v-card-title class="headline"> 使い方 </v-card-title>
        <v-card-text>
          <h4>Jantama-Analyzerは雀魂のプレイ画像を解析し、手牌の受け入れ枚数、点数期待値、和了確率、聴牌確率を計算します。</h4>
          <br />
          <h4>トップページの「Gyazoにログイン」からGyazoにログインするとユーザーのキャプチャ一覧が表示されるので、解析したいプレイ画像を選択します。
            選択したページにある「画像の解析」ボタンを押し、しばらくすると解析結果が表示されます。</h4>
          <hr class="my-3" />
          <h3>キャプチャ一覧</h3>
          <v-img
            :src="require('@/assets/howto/list.jpg')"
          ></v-img>
          <br />
          <h3>プレイ画像</h3>
          <v-img
            :src="require('@/assets/howto/image.jpg')"
          ></v-img>
          <br />
          <h3>解析結果</h3>
          <v-img
            :src="require('@/assets/howto/analyze.jpg')"
          ></v-img>
          <hr class="my-3" />
          <h2>注意点</h2>
          <br />
          <ul>
            <li>雀魂はPCのブラウザ版のプレイを想定しています。</li>
            <li>プレイ画像はGyazoのブラウザ拡張機能を用いたページ全体のスクリーンショットを想定しています。</li>
            <li>それ以外の部分的に切り抜いた画像やアスペクト比が異なる画像では手牌の情報を上手く検出できない可能性があります。</li>
            <li>パフォーマンス向上のため、Gyazoのメタ情報に「雀魂」の記載がないキャプチャ画像は「画像の解析」を行うことが出来ません。</li>
            <li>Gyazoの仕様上、freeユーザーは直近10枚までの画像しか表示されません。</li>
            <li>副露は考慮しません。門前手以外のプレイ画像には現在対応していません。</li>
            <li>七対子手・国士無双手には現在対応していません。</li>
            <li>ドラの白は認識精度が甘いです、ご了承ください。</li>
          </ul>
          <hr class="my-3" />
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn color="accent" to="/"> 戻る </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>
