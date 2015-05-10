<?php
    //$conn = mysql_connect('127.0.0.1:2350', 'root', 'root') or die(mysql_error());
    /*$conn = mysql_connect('10.1.72.77:8075', 'root', 'root') or die(mysql_error());
    mysql_select_db('vid2');
    mysql_query("SET NAMES UTF8");

    function error($sql){
        echo "SQL: $sql\n\nERROR: " . mysql_errno() . " - ". mysql_error() . "\n\n";
    }*/
    function http_post($url, $params, $use_zip = TRUE) {
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $url);
        curl_setopt($ch, CURLOPT_HEADER, $use_zip);
        curl_setopt($ch, CURLOPT_BINARYTRANSFER, 1);
        curl_setopt($ch, CURLOPT_POST, 1);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $params);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($ch, CURLOPT_HTTPHEADER, array(
            'Accept:text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
            'Accept-Encoding: gzip',
            'Connection:keep-alive',
            'Cache-Control:max-age=0'
        ));
        $result = curl_exec($ch);

        if ($use_zip) {
            $pos = strpos($result, "\r\n\r\n");
            $head = substr($result, 0, $pos);
            $body = substr($result, $pos + strlen("\r\n\r\n"));
            if (strpos($head, 'gzip') > 0) {
                $body = gzinflate(substr($body, 10, -8));
                                                                     }
        } else {
          $body = $result;
        }
        return $body;
    } 
    function addVideo($data){
        $json = json_encode($data);
        echo http_post('http://api.tv.uc.cn:8810/v3/video/formal?from=zlx&sign=37bb6553d5e93db013935fcc0f6f2b87', 'json='.urlencode($json), FALSE), "\n";
        //var_dump($data);
    }
    function main(){
        $page = 1;
        $page_size = 100;
        $start_time = time() - 3600;
        while(true){
        	$json_url = "http://www.1yingyuan.com/api/movie.json?play_type=qvod&start_time={$start_time}&page={$page}&page_size={$page_size}";
        	echo "fetch:", $json_url, "\n";
            $json = file_get_contents($json_url);
            $json = json_decode($json, true);
            if($json['Status'] == 200){
            	foreach($json['Data']['Videos'] as $video){
                    if(count($video['Items']) === 0/* || $video['Genres'] === '纪录'*/) continue;
            		$item = $video['Items'][0];
            		$sn = $item['Sn'];
            		$url = "http://www.1yingyuan.com/detail/{$video['Id']}.html";
            		$data = array(
                        'channel' => $video['Channel'],
                        'title' => $video['Title'],
                        'summary' => $video['Summary'],
                        'actors' => $video['Actors'],
                        'directors' => $video['Directors'],
                        'genres' => $video['Genres'],
                        'update_time' => $video['UpdateTime'],
                        'source' => '快播',
                        'image_small' => $video['Poster'],
                        'op' => ($video['Status'] === 'NORMAL') ? 'ADD' : 'DEL',
                        'items' => array(
                            array(
                            	'sn' => 1,
                            	'platform' => 3,
                            	'title' => $video['Title'],
                            	'update_time' => $item['CreateTime'],
                                'playurl' => $url
                            )
                        )
                    );
                    if($video['Year'] > 0){
                        $data['year'] = $video['Year'];
                    }
                    if(!empty($video['Alias'])){
                        $data['alias'] = $video['Alias'];
                    }
                    addVideo($data);
            	}
            	if(count($json['Data']['Videos']) === 0){
            		break;
            	}
                $page++;
            }
        }
    }
    main();
