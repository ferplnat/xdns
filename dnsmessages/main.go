package dnsmessages

func handleDNSRequest(buf []byte){
    headerLength := int(buf[2]&0xF) * 4

    questionStartPosition := headerLength
    numOfQuestions := int(buf[4]<<8) | int(buf[5])

    for i := 0; i < numOfQuestions; i++ {
        question, questionLength := parseQuestion(buf[questionStartPosition:])        
        questionStartPosition += questionLength
    }
}

func parseQuestion(questionData []byte) ([]byte, int) {
    question := make([]byte, 0)
    position := 0

    for questionData[position] != 0 {
        if position != 0 {
            question = append(question, '.')
        }

        if (questionData[position] & 0xC0) == 0xC0 {
            // Handle compression, if needed.
            offset := int(questionData[position]&0x3F)<<8 | int(questionData[position+1])
            position = offset
        } else {
            labelLength := int(questionData[position])
            label := questionData[position+1 : position+1+labelLength]
            question = append(question, label...)
            position += labelLength + 1
        }
    }

    position++
    return question, position
}
