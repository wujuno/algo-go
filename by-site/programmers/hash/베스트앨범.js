function solution(genres, plays) {
    const genrePlayCount = {}; // 장르별 총 재생 횟수
    const songsByGenre = {};   // 장르별 곡 정보 목록

    // 1. 데이터 수집
    genres.forEach((genre, i) => {
        const play = plays[i];

        // 총 재생 횟수 합산
        genrePlayCount[genre] = (genrePlayCount[genre] || 0) + play;

        // 장르별 곡 리스트에 저장 (id와 play 횟수)
        if (!songsByGenre[genre]) {
            songsByGenre[genre] = [];
        }
        songsByGenre[genre].push({ id: i, play: play });
    });

    // 2. 장르 순위 정하기 (총 재생 횟수 내림차순)
    const sortedGenres = Object.keys(genrePlayCount).sort((a, b) => {
        return genrePlayCount[b] - genrePlayCount[a];
    });

    // 3. 베스트 앨범 구성
    const result = [];

    sortedGenres.forEach(genre => {
        // 해당 장르의 곡들을 정렬 (재생 횟수 내림차순 -> 고유 번호 오름차순)
        const songs = songsByGenre[genre].sort((a, b) => {
            if (a.play === b.play) {
                return a.id - b.id; // 재생 횟수 같으면 id 낮은 순
            }
            return b.play - a.play; // 기본은 재생 횟수 높은 순
        });

        // 최대 2곡까지만 결과에 추가
        result.push(...songs.slice(0, 2).map(song => song.id));
    });

    return result;
}