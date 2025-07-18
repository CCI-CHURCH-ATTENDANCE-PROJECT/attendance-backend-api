package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/skip2/go-qrcode"
)

// GenerateQRCode generates a QR code for the given data
func GenerateQRCode(data string, size int) (string, error) {
	// Generate QR code
	qr, err := qrcode.New(data, qrcode.Medium)
	if err != nil {
		return "", fmt.Errorf("failed to create QR code: %w", err)
	}

	// Convert to PNG bytes
	pngBytes, err := qr.PNG(size)
	if err != nil {
		return "", fmt.Errorf("failed to convert QR code to PNG: %w", err)
	}

	// Encode to base64
	base64Image := base64.StdEncoding.EncodeToString(pngBytes)
	return base64Image, nil
}

// GenerateQRCodeBuffer generates a QR code and returns as buffer
func GenerateQRCodeBuffer(data string, size int) (*bytes.Buffer, error) {
	qr, err := qrcode.New(data, qrcode.Medium)
	if err != nil {
		return nil, fmt.Errorf("failed to create QR code: %w", err)
	}

	pngBytes, err := qr.PNG(size)
	if err != nil {
		return nil, fmt.Errorf("failed to convert QR code to PNG: %w", err)
	}

	return bytes.NewBuffer(pngBytes), nil
}
