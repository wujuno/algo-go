function solution(record) {
    const answer = [];
    const nicknames = {}; // { "uid1234": "Prodo", ... } 형태로 저장될 객체

    // 1단계: 최종 닉네임 정보를 업데이트합니다.
    for (const line of record) {
        const [cmd, uid, nick] = line.split(' ');

        // Enter이거나 Change일 때만 닉네임 정보를 갱신합니다.
        // Leave일 때는 nick이 undefined이므로 체크가 필요합니다.
        if (cmd === 'Enter' || cmd === 'Change') {
            nicknames[uid] = nick;
        }
    }

    // 2단계: 기록을 다시 돌며 메시지를 생성합니다.
    for (const line of record) {
        const [cmd, uid] = line.split(' ');

        if (cmd === 'Enter') {
            answer.push(`${nicknames[uid]}님이 들어왔습니다.`);
        } else if (cmd === 'Leave') {
            answer.push(`${nicknames[uid]}님이 나갔습니다.`);
        }
        // Change는 메시지를 남기지 않으므로 무시합니다.
    }

    return answer;
}