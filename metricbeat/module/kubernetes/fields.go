// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package kubernetes

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

// AssetKubernetes returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/kubernetes.
func AssetKubernetes() string {
	return "eJzsXc9v47byv+evIPaUfpH68MXDO+TwgDZt8YLdboP86B4eHgxaGttsJFJLUsm6f/0DqZ+WSIqyaCebSIeia8czH84Mh8PRcPgjeoTdJXrMV8ApSBBnCEkiE7hEHz7WH344QygGEXGSScLoJfrXGUIINX+AUpCcROrXHBLAAi7RBp8hJEBKQjfiEv3ngxDJhwv0YStl9uG/6rst43IZMbomm0u0xomAM4TWBJJYXGoGPyKKU+jAU4/cZYoDZ3lWfmKAp55rumY8xepjhGmMhMSSCEkigdgaZSwWKMUUbyBGq12Lz6Kk0EbTRoQzIoA/Aa+/MYFyAOvI76eba1QQbImyevZFWj1daG14HL7mIOQiSghQufcnFc5H2D0zHne+c6BVz5Wmh+AbRLnSa8VIOFFwECznEYTDcVtQhhgZaXcBiHx1TAw28j0YEcvCA0CaLDqPklxI4BeaqchwBBe1dH5w4noCvgoH69/39zeoR7JnmSwOKArNs0eyz5NKoHKpGIVXQ4lBs0A9Fl0sMd8teU7DwfgCcgscyS1UPFAuQKCY71CXURfMI6FdbhOQfCQ0Vt61pD6gkjRjNKyPqkiiLaZxorxUSyhONF3fPRGJcuqaJFqzSjMebuIJuCAsoGmUBGsU/WF2IWjJ7S1uEyFUk8REuMs8BbllAe1RT0wD0d6gmQhohvWIu1QrthlnEQhh5GgyRNN636YXZflCQNT7vqIZs3yVdP1ebyBXNw9IQMRo3EXWcEohZXynlnUSA5WL1a6JzPp8E0Y3hi+LuOwS2X68h+pn9UeIUFTxLDEMQXwiXOY4OSXCkuUQwHUsFiwDuohY3vN+g9D2WH/O0xVw5XEVQbQmCdR/wLhdjUJiLiEOYDR3hcEgQWgE2sWUxl3xME4AtREIZv31uppzHe0vcrHIgEdAJUlg8X/WEbLVXxCZFFB8sRwjh2rOVyBQSiLOyumEGjh2nZiGIfJ0on7cuKI8zRMsyRMgEysXtOnGW0HTlPQKVdEfBCLI31DM7JCaHgNaIRil1hZkl1ZDOKQ9jCNV3IJ5DA0r8g4MImNUwIuqt4AwRr990MdXcBult4b7QEOouIRiJtUP+sPbVDUw40pTpEEWLv5W3paltkp8ICyQIcvSGXK4IC9gtGDM3bSZJVgCjXaHWLJJW6IieKFMVCEo/k2KwKm9Jg1CCmdCNSY6XjCrPHoEedIlp2SNtkRItuE4RQUIO1jfUGIMiopmoUlf5R0ncmiw0HYgXHzoB+YF9Nig9tdklHOu/Nh02V3TdUI2W+lh6oxueE4poZugW5XGf0Z60VK/RiUjd1YZZBQvCrkH8eRN0r/UpkBYai5G9jiPiVzAk00RY9lrekjTM4+3YMhBQYM4IM+KZJd5s9ZQiQmd9o6jJd2aXpBXHHpnuZQkNadyYyy7XwwkbO4UQdQj2EqveK/iQxnKmweUC7wBgyBsw25D0b+1zkMTIBfVvUEybiI8THyIQZuJwSl32Vh8SfUMyLf9XNVGp6R+xTiUoqeYWhesPbSYMiUWG+hBwJ5gC6OAeIBhDYvFsMiMa1KDSkQ4gXi5Thi2/WG15Sh3OWbDGj0MJWAsEG6RVX5HbgFJJnGi4SOcJCzCEq8SUL9zjjchKZHf5YBjWBMKcTGCOgffOMNz9YlVKIisUU71byE2v8ZL2MY/izwwqk9so4LxNRvpkvATJgk2p6KmuyXbfhj5zL+hTTXy17WWTj1UFOEMR0TuVABspl771fIv3750Ckv2l4xyeW9fKtqx+wuFKE9gf18xbYU3x/Ao4FJ2r22gmSfW4bReh3BwBx6hUClGPoAsdhkekDYNA6D9N1nBEkjvw1F3LXDgZdzxAurXJZBCDNbhvv7o8vfWAEYGmBYTQN9DjOkz7AlhZmkW9kizLSTeK1lAb2qm3N7duedJBfiZ8UdCNwLsKbG3II8vxTCRAOknlwxvYI3zxJBUHJMqNCNqcliKDbLwqddO/BfjJ8KjeVlR1bOHMbkOWPPzHvYVt4xJXdUidkJCOnqL8T5CHrOU2kH4e9+JmSVUxt8vtyM7wU7jwbDHaGf5OUsS4MVBiEnZ/quaWHmswpnrX4H0zfa/SEHqKWvUT130euJiV/XfcOw+4xT8aqr/ZjQg32u65lhInkcy59AnPpf2FsOZS3vn0t65tNdjGHNprxnIXNrrjXEu7Z1Le+fS3umlvYYoc2yx7zPjj19zyM0R5yFLnwINKuAsCvCmL+efCoJ1pV25mLtiiZyuCSViGySceKiJ+bDGcRzChr9UelEEBww5hkxug/LUFAenj+QkyHxt+LbrmTV188aMxbCI1JY9ksy8vz7EcOGJRDqSCBkD6xcXFWWXwW4BJ3Ibokq8YV5TReZU0DEq9N2cCjyW91X+7G72XiXZB1n7JMAx8AURyxQLacnJrBhLAHcDvaEj7NvmDLvWNRGow+Osi0bXrp512Y9IWd1vod2Io6iFrbJWoNYhPTfqb+QWS4Q5oA1Q4FgWnUOqyuHSr+5xIFRtbJVwP3b7mKARpa92A7Po2intq2J5VVwQh4jxWBRyr41PkhSKzzLMJYnyBPNCCGiLBWKRLkePDQj1LyVOMwPKvjNxpf3WhAu5LFlRS/eO8aW+9xVANU7NAzU81Gddq2of/Tg6IMViAE+TCxG9d3H2/JYTxO8FqdIYIG66BZAnoAaJRCzbLSUzgWiWNSw6u70D0d1qSr7gakPsduE4kPv9Lqvfs7s5piBxjPfS2nbLH9BHQQlhIVhEtKN5JnLr1IlrLpln5fhFvvZDHHA3/4Ncc8DjdcXePNAMCKNuyR81x1xydvPUvXbCMtYkEaHoeUuibel1n7FoFh0jmioVvgzeQuTPsoVIWyDuzHtOAr7NeKDkaw5I54fJmqgYgbWAGPIDdSYUkvUyIfQxIJjbT4hDxkEoNGV7GZtDIPSJJU8QLw0Yj+UXKp4mubg8BM5IeMv56ea6bkBTWo9DXWE7ESnej2U3ogHGYZ0HbTkPB9PjzdeK8gjRh52wD9e/DPBu70CnBPCtE2Z60zAfLpsPl1me0IfLPit7+77Plc1l5qZnLjPvPOHKzOc64g7guY7YDHyuI3bUEVOQym6C+Wv+7U0b3y1EQJ50qtZGq04oc256JeWJ2RfPNxufOlvzthVyzzEVKZHy9ejk3qiTOhM9F+0Xj6c0f5vr9UcKaC7Vb56ecN5DlX7rXbPlTHAX1CkOczeoXscx7gaP7Sh3HdPk1JrBOcRvk1RFgEc6lm9fE4YZDDFBnjMc+aZIfGY6GpdKuU51xDt+1UCeKwd6z2L0WFvQGGf3DkVoXoHqzereMZspOeyMxd9lCnvekRbPvCNtnu9JId/djvRdvDN6JW9JerBeaX+UMd333lvHPbWw1u1QRLcfil+rvcBvyeYXQh3Yr3VezX2Hgk62g5sPvY/U4N6ksQ+58wJx+dbfIBZiee69R7RvId74K+ZCIPW5eSURfWBwQCwZ3sDyaG8yC1Deb1WXp0Bjf6fa6tjwbTdlB986QqJpedzmOtj+pD7oY+hRcnAdva3tSZNTjoPUzJvanbSq5Lt9SaZw6ZGrBddtRzJVavv0XO0+xpx2GW704Txi6dnkY1yLD8fEc3uwQ5p7jGrtERiZs6mHZ0sPB6QJ7Tx8mnn4G8aYRh7WNh6HWfXoBh7OU/8+zTuCtO4Y27gjFCLnef/xLTt8jdO7XcehzTr8teoPdqCNw8gmHWFci397jtHNOQ7XpaExx8FtOcIq0q8hx9h2HKFU6d2IY3wbjtEiMpHxa8BxkN2YgsOhbhuHHEX27LNRL4c7GnktSk6mj/kKikC9DNd3NDJmvweWtjwB4bkyDIv/bkejGwXnVpHt3LLG1vUHQ/fl2dFNMw8rPo+b1+yYrLevhfQzVuhD16913m9mXP9xSugmmNo/F6RRi/aoG/Y8IU6MXZ0gRxjAAMqTWIN7MHaT6OUNRLSFOE+mdU1t5Q5qenPi4B0kDnqHTg9kM9QPtRWb5EmQgd2VdoqwlJBmsk+64ln7g4Bs1XQ10Z0TMnNCZgjSnJCZEzIjEc0JmTkhMydk5oTMnJAxYnA2BCz4m9oBOiGMaQXY2411G/AdtkjC/8PpN6a/0hhJhoDGrcGYlyVP2FMSEyPQOCZgF9G0GWHG5JqJGYsXGQe1TVEIdP/QdCqMGxajhigqiToQlBulEHwrUs5R1xIvFXTKAO/OYCzDK0kP8bSYzgTCa8Ho4ZiYNLVZ6VmX8eu+cP/g9lA98TR3shNzHzshsczDHbjOtljYKwbNA+gOwlWPXA9HM0LnZRPYC/SMidT/I4GnhGL3xYqAY/uZcHNDXU+UDULNxCzfvYhJ7UDt9ViEStj0Ov8eAKbgM9gcu9dFtA1mkv6+FBpC5zWqK911UintimOx/cRY9jOOHtl6fYF+5VyfDrvJk+QC1f9bft9XrXoYr7WvPND5FUuzBCTEF40krjClTN7mVLNg/AL98cfvH0mSQPxDOfyFcaKMOQMy2GteVyHbzj4UdG3Fx6PUfnXzoHuBiYKlQ+9VUHsSSCU7iJGZ4b6cXOdEBooWMw6RcgWX6J+Lf4RAXmPxFKgL+zC8qSWZNqmftD9ZocTj3x01JIKyyLsonh/sb1Ap8OVxN2qr6vdtJ2NjyBK2Syd2aW9FNQ3BIGFNhg2nnV1r7oCcPhqRFlxMC2+j2ywhEfaPeg7CUXFBhK7ZyDkTgyDc0XhkUlTwS4OxeZdccmxQn4sMoinHQ0JhbA7UW/TWOuRFTwerxcsDWBYbG3gHB1Xw6QNqn/oK5BxCdoANuQVxB/eTQljdhLQd16NzyXO4QGucCFDxZ04fKXum9nmT0zKV4jTSSVsQjXKPj8sZhoxrW4fNjhdK1p1g20fb3HFk1dpkANSEnoIVprqJyunav7Zk/lLBymfbScOhKKtWzIsiL9G6G+CY24+gYLrTZzaPZZpt3WTMcSltrZCjwtFHVrttmF5vfxdivsej97F/mKiQXd8YmW2ZkMvjcFSkbWxHLsLjGJeL5WEdL46YUuzALHOKt1VO8QZoTOhmsVgcmkoMiW5a3FFVO9pj0JBYa24mvBd9tN2dGYTawZYEyyPM013BEbeObaj2PWyIfvATzvhu9y87LPeKGXB0W/zjznAy3ndX+1K43HM4HCo1f8diYyt9uvlYQisvOdJXh5Sc0GqnawEacPrtFmdJt765jTPBK3B5l1BSXOdJsqu4DUqzvbrBOk/COZaKYjDPYr7jyyq8AcEpodWXetXXkaFzyFi0/UFXRd2VI+ha3wlc3Z7wah0e5O2OPD9ahy2q6bFnczYhohdwe70UngtgBa5xAMfWc8vVEFps3fpcX1jdtZJbYF+HmivlegBrTu8JCWkof1c06Gw1Ggri9AxlGGhEvYOpwV794qR35GOvqHy+csoPkIvq3iDf3JVT821TBnJzH712iPTGu17NFyvtP/PFSn54hpuAPbEkT0O9iSyIBQlIejHDpFjkzwKYNRCZb7opH8/5N990M1ZA8003zfMub7p58Lzf5gTXyfxmuUSmC+UUV+0UQV4J5n8BAAD//x4OEII="
}
