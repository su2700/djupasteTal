package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    texttospeech "cloud.google.com/go/texttospeech/apiv1"
    texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

func main() {
    ctx := context.Background()

    // Creates a client.
    client, err := texttospeech.NewClient(ctx)
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }
    defer client.Close()

    // The text to synthesize.
    text := "Hej, hur mår du?"

    // Perform the text-to-speech request on the text input with the selected
    // voice parameters and audio file type.
    req := &texttospeechpb.SynthesizeSpeechRequest{
        Input: &texttospeechpb.SynthesisInput{
            InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
        },
        Voice: &texttospeechpb.VoiceSelectionParams{
            LanguageCode: "sv-SE",
            Name:         "sv-SE-Standard-A",
        },
        AudioConfig: &texttospeechpb.AudioConfig{
            AudioEncoding: texttospeechpb.AudioEncoding_MP3,
        },
    }

    resp, err := client.SynthesizeSpeech(ctx, req)
    if err != nil {
        log.Fatalf("Failed to synthesize speech: %v", err)
    }

    // The response's audioContent is binary.
    err = ioutil.WriteFile("output.mp3", resp.AudioContent, 0644)
    if err != nil {
        log.Fatalf("Failed to write audio file: %v", err)
    }

    fmt.Println("Audio content written to file: output.mp3")
}
